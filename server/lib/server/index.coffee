# TODO: we have to move kd related functions to somewhere else...

{argv} = require 'optimist'
Object.defineProperty global, 'KONFIG',
  value: require('koding-config-manager').load("main.#{argv.c}")

{
  webserver
  mongo
  mq
  projectRoot
  kites
  uploads
  basicAuth
  neo4j
  github
}       = KONFIG
webPort = argv.p ? webserver.port
koding  = require './bongo'

processMonitor = (require 'processes-monitor').start
  name                : "webServer on port #{webPort}"
  stats_id            : "webserver." + process.pid
  interval            : 30000
  librato             : KONFIG.librato
  limit_hard          :
    memory            : 300
    callback          : ->
      console.log "[WEBSERVER #{webPort}] Using excessive memory, exiting."
      process.exit()
  die                 :
    after             : "non-overlapping, random, 3 digits prime-number of minutes"
    middleware        : (name,callback) -> koding.disconnect callback
    middlewareTimeout : 5000

_          = require 'underscore'
async      = require 'async'
{extend}   = require 'underscore'
express    = require 'express'
Broker     = require 'broker'
fs         = require 'fs'
hat        = require 'hat'
nodePath   = require 'path'
http       = require "https"
{JSession} = koding.models
app        = express()

{
  error_
  error_404
  error_500
  authTemplate
  authCheckKey
  authenticationFailed
  findUsernameFromKey
  findUsernameFromSession
  fetchJAccountByKiteUserNameAndKey
  renderLoginTemplate
  serve
  isLoggedIn
  saveOauthToSession
  getAlias
}          = require './helpers'


# this is a hack so express won't write the multipart to /tmp
#delete express.bodyParser.parse['multipart/form-data']

app.configure ->
  app.set 'case sensitive routing', on
  app.use express.cookieParser()
  app.use express.session {"secret":"foo"}
  app.use express.bodyParser()
  app.use express.compress()
  app.use express.static "#{projectRoot}/website/"

app.use (req, res, next)->
  res.removeHeader "X-Powered-By"
  next()

if basicAuth
  app.use express.basicAuth basicAuth.username, basicAuth.password

process.on 'uncaughtException',(err)->
  console.log 'there was an uncaught exception'
  console.log process.pid
  console.error err

app.use (req, res, next) ->
  {JSession} = koding.models
  {clientId} = req.cookies
  clientIPAddress = req.connection.remoteAddress
  res.cookie "clientIPAddress", clientIPAddress, { maxAge: 900000, httpOnly: false }
  JSession.updateClientIP clientId, clientIPAddress, (err)->
    if err then console.log err
    next()

app.get "/-/auth/check/:key", (req, res)->
  {key} = req.params

  console.log "checking for key"
  authCheckKey key, (ok, result) ->
    if not ok
      console.log "key is valid: #{result}" #keep up the result to us
      res.send 401, authTemplate "Key is not valid: '#{key}'"
      return

    {JKodingKey} = koding.models
    JKodingKey.fetchKey
      key     : key
    , (err, kodingKey)=>
      if err or not kodingKey
        res.send 401, authTemplate "Key doesn't exist"
        return

      res.send 200, {result: 'key is added successfully'}

app.get "/-/auth/register/:hostname/:key", (req, res)->
  {key, hostname} = req.params

  authCheckKey key, (ok, result) ->
    if not ok
      console.log "key is not valid: #{result}" #keep up the result to us
      res.send 401, authTemplate "Key is not valid: '#{key}'"
      return

    isLoggedIn req, res, (err, loggedIn, account)->
      if err
        # console.log "isLoggedIn error", err
        res.send 401, authTemplate "Koding Auth Error - 1"
        return

      if not loggedIn
        res.send 401, authTemplate "You are not logged in! Please log in with your Koding username and password"
        return

      findUsernameFromSession req, res, (err, notUsed, username) ->
        if err
          # console.log "findUsernameFromSession error", err
          res.send 401, authTemplate "Koding Auth Error - 2"
          return

        if not username
          res.send 401, authTemplate "Username is not defined: '#{username}'"
          return

        console.log "CREATING KEY WITH HOSTNAME: #{hostname} and KEY: #{key}"
        {JKodingKey} = koding.models
        JKodingKey.fetchByUserKey
          username: username
          key     : key
        , (err, kodingKey)=>
          if err or not kodingKey
            JKodingKey.createKeyByUser
              username : username
              hostname : hostname
              key      : key
            , (err, data) =>
              if err or not data
                # console.log "createKeyByUser error", key, err
                res.send 401, authTemplate "Koding Auth Error - 3"
              else
                res.send 200, authTemplate "Authentication is successfull! Using id: #{hostname}", key, hostname

          else
            res.send 200, authTemplate "Authentication already established!"


app.get "/-/imageProxy", (req, res)->
  if req.query.url
    require('request')(req.query.url).pipe(res)
  else
    res.send 404

s3 = require('./s3') uploads.s3, findUsernameFromKey
app.post "/-/kd/upload", s3..., (req, res)->
  {JUserKite} = koding.models
  for own key, file of req.files
    console.log "--------------------------------->>>>>>>>>>", req.account
    zipurl = "#{uploads.distribution}#{file.path}"
    JUserKite.fetchOrCreate
      kitename      : file.filename
      latest_s3url  : zipurl
      account_id    : req.account._id
      hash: req.fields.hash
    , (err, userkite)->
      if err
        console.log "error", err
        return res.send err
      userkite.newVersion (err)->
        if not err
          res.send {url:zipurl, version: userkite.latest_version, hash:req.fields.hash}
        else
          res.send err

app.get "/Logout", (req, res)->
  res.clearCookie 'clientId'
  res.redirect 302, '/'

if uploads?.enableStreamingUploads

  s3 = require('./s3') uploads.s3, findUsernameFromSession

  app.post '/Upload', s3..., (req, res)->
    res.send(for own key, file of req.files
      filename  : file.filename
      resource  : nodePath.join uploads.distribution, file.path
    )

  # app.get '/Upload/test', (req, res)->
  #   res.send \
  #     """
  #     <script>
  #       function submitForm(form) {
  #         var file, fld;
  #         input = document.getElementById('image');
  #         file = input.files[0];
  #         fld = document.createElement('input');
  #         fld.hidden = true;
  #         fld.name = input.name + '-size';
  #         fld.value = file.size;
  #         form.appendChild(fld);
  #         return true;
  #       }
  #     </script>
  #     <form method="post" action="/upload" enctype="multipart/form-data" onsubmit="return submitForm(this)">
  #       <p>Title: <input type="text" name="title" /></p>
  #       <p>Image: <input type="file" name="image" id="image" /></p>
  #       <p><input type="submit" value="Upload" /></p>
  #     </form>
  #     """

app.get "/-/presence/:service", (req, res) ->
  # if services[service] and services[service].count > 0
  res.send 200
  # else
    # res.send 404

app.get '/-/services/:service', require './services-presence'

app.get "/-/status/:event/:kiteName",(req,res)->
  # req.params.data

  obj =
    processName : req.params.kiteName
    # processId   : KONFIG.crypto.decrypt req.params.encryptedPid

  koding.mq.emit 'public-status', req.params.event, obj
  res.send "got it."

app.get "/-/api/user/:username/flags/:flag", (req, res)->
  {username, flag} = req.params
  {JAccount}       = koding.models
  JAccount.one "profile.nickname" : username, (err, account)->
    if err or not account
      state = false
    else
      state = account.checkFlag('super-admin') or account.checkFlag(flag)
    res.end "#{state}"

app.post "/-/oauth/:provider/callback", (req,res)->

app.get "/-/oauth/:provider/callback", (req,res)->
  {provider} = req.params
  code = req.query.code
  access_token = null

  unless code
    {loginFailureTemplate} = require './staticpages'
    serve loginFailureTemplate, res
    return

  headers =
    "Accept"     : "application/json"
    "User-Agent" : "Koding"

  authorizeUser = (authUserResp)->
    rawResp = ""
    authUserResp.on "data", (chunk) -> rawResp += chunk
    authUserResp.on "end", ->
      {access_token} = JSON.parse rawResp
      if access_token
        options =
          host    : "api.github.com"
          path    : "/user?access_token=#{access_token}"
          method  : "GET"
          headers : headers
        r = http.request options, fetchUserInfo
        r.end()

  fetchUserInfo = (userInfoResp) ->
    rawResp = ""
    userInfoResp.on "data", (chunk) -> rawResp += chunk
    userInfoResp.on "end", ->
      {login, id, email, name} = JSON.parse rawResp
      if name
        [firstName, restOfNames...] = name.split ' '
        lastName = restOfNames.join ' '

      {clientId} = req.cookies
      resp = {provider, firstName, lastName, login, id, email, access_token,
              clientId}

      if not email? or email is ""
        options =
          host    : "api.github.com"
          path    : "/user/emails?access_token=#{access_token}"
          method  : "GET"
          headers : headers
        r = http.request options, (newResp)-> fetchUserEmail newResp, resp
        r.end()
      else
        renderLoginTemplate resp, res

  fetchUserEmail = (userEmailResp, originalResp)->
    rawResp = ""
    userEmailResp.on "data", (chunk) -> rawResp += chunk
    userEmailResp.on "end", ->
      email = JSON.parse(rawResp)[0]
      originalResp.email = email
      renderLoginTemplate originalResp, res

  options =
    host   : "github.com"
    path   : "/login/oauth/access_token?client_id=#{github.clientId}&client_secret=#{github.clientSecret}&code=#{code}"
    method : "POST"
    headers : headers
  r = http.request options, authorizeUser
  r.end()

app.get '/:name/:section?*', (req, res, next)->
  {JName} = koding.models
  {name} = req.params
  return res.redirect 302, req.url.substring 7  if name in ['koding', 'guests']
  [firstLetter] = name
  if firstLetter.toUpperCase() is firstLetter
    next()
  else
    {JGroup} = koding.models
    isLoggedIn req, res, (err, loggedIn, account)->
      JName.fetchModels name, (err, models)->
        if err then next err
        else unless models? then res.send 404, error_404()
        else
          models.last.fetchHomepageView account, (err, view)->
            if err then next err
            else if view? then res.send view
            else res.send 500, error_500()

app.get "/", (req, res)->
  if frag = req.query._escaped_fragment_?
    res.send 'this is crawlable content'
  else
    {JGroup} = koding.models
    isLoggedIn req, res, (err, loggedIn, account)->
      if err
        res.send 500, error_500()
        console.error err
      else if loggedIn
        # go to koding activity
        activityPage = JGroup.renderKodingHomeLoggedIn {account}
        serve activityPage, res
      else
        # go to koding home
        homePage = JGroup.renderKodingHomeLoggedOut()
        serve homePage, res

app.get '*', (req,res)->
  {url}            = req
  queryIndex       = url.indexOf '?'
  [urlOnly, query] =\
    if ~queryIndex
    then [url.slice(0, queryIndex), url.slice(queryIndex)]
    else [url, '']

  alias      = getAlias urlOnly
  redirectTo =\
    if alias
    then "#{alias}#{query}"
    else "/#!#{urlOnly}#{query}"

  res.header 'Location', redirectTo
  res.send 302

app.listen webPort
console.log '[WEBSERVER] running', "http://localhost:#{webPort} pid:#{process.pid}"
