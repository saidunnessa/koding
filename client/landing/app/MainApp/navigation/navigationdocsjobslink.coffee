class NavigationDocsJobsLink extends KDCustomHTMLView

  constructor:(options = {}, data)->

    options.tagName    = "span"
    options.cssClass   = "title"

    super options, data

    @icon      = new KDCustomHTMLView
      tagName  : "span"
      cssClass : "main-nav-icon #{__utils.slugify @getData().title}"

    @docsLink  = new KDCustomHTMLView
      tagName    : "a"
      partial    : "Docs"
      cssClass   : "ext"
      attributes :
        href     : "http://koding.github.io/docs/"
        target   : "_blank"

    @jobsLink  = new KDCustomHTMLView
      tagName    : "a"
      partial    : "Jobs"
      cssClass   : "ext"
      attributes :
        href     : "http://koding.github.io/jobs/"
        target   : "_blank"

  click:(event)->

  viewAppended: JView::viewAppended

  pistachio: ->
    """
    {{> @icon}} {{> @docsLink}} / {{> @jobsLink}}
    """
