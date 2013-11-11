myRealEscapeString = (str) ->
  return str.replace(/[\0\x08\x09\x1a\n\r"'\\\%]/g, (char) ->
    switch char
      when "\0" then "\\0"
      when "\x08" then "\\b"
      when "\x09" then "\\t"
      when "\x1a" then "\\z"
      when "\n" then "\\n"
      when "\r" then "\\r"
      when "\"" then "\\" + char
      when "'" then "\\" + char
      when "\\" then "\\" + char
      when "%" then "\\" + char
  )

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

  m.score = ko.computed ->
    score = 0
    for flag in m.flags()
      score += flag.value
    score

  m

scoreboardModel = ->
  validateFlagWthServer = (publicId, flagHash, cb) ->
    console.dir publicId
    console.dir flagHash
    cb("response")

  m =
    users: ko.observableArray []

  m.load = -> loadModel m

  m.validateFlag = (d, e) ->
    publicId = myRealEscapeString($("#publicId").val())
    flagHash = myRealEscapeString($("#flagHash").val())
    validateFlagWthServer publicId, flagHash, (response) ->
      console.log response

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
