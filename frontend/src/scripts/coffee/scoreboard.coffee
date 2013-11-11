server = ->
  url = (route) -> "http://localhost:9999/#{route}"
  
  call = (route, callback) ->
    $.getJSON route, {}, callback

  submit = (type, url, cb, data) ->
    $.ajax
      url: url
      data: JSON.stringify (data || {})
      type: type
      dataType: "json"
      success: if cb? then cb else ->

  m =
    fetch: (route, cb) -> call(url(route), cb)
    update: (route, cb, data) -> submit("PUT", url(route), cb, data)
  m

loadModel = (model) ->
  model.users []
  server.fetch "scoreboard", ({scoreboard: data}) ->
    for score in data.scores
      model.users.push userModel(score)

userModel = (d) ->
  m =
    id: d.user.public_id
    flags: ko.observableArray d.flags

  m.flagsFound = ko.computed ->
    m.flags().length

  m

scoreboardModel = ->
  m =
    users: ko.observableArray []

  m.load = -> loadModel m

  m.toggleSubmission = ->
    $("#FlagSubmission").toggle("slide", { direction: "up" } )

  m.totalFlagsFound = ko.computed ->
    found = 0
    for user in m.users()
      found += user.flagsFound()
    found

  m

scoreboardViewModel = () ->
  model = scoreboardModel()
  model.load()

  $('#Counter').countdown({
    image: '../../images/digits.png',
    startTime: '12:12:00'
  });

  model

$(document).ready(->
  server = new server()
  ko.applyBindings new scoreboardViewModel()
)
