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
      contentType: "application/json"
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

reloadModel = (model) ->
  server.fetch "scoreboard", ({scoreboard: data}) ->
    for score in data.scores
      userMatch = ko.utils.arrayFirst model.users(), (item) ->
        score.user.public_id == item.id
      if userMatch?
        userMatch.flags score.flags

createUserFlagModel = (data) ->
  user:
    public_id: data.public_id
  flags:
    [
      { tag: data.tag, value: data.value }
    ]

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
    userflagModel = createUserFlagModel { public_id: publicId, tag: flagHash, value: 1 }
    server.update "validateFlag", cb, userflagModel

  m =
    users: ko.observableArray []

  m.load = -> loadModel m

  m.reload = -> reloadModel m

  m.validateFlag = (d, e) ->
    publicId = parseInt(myRealEscapeString($("#publicId").val()))
    flagHash = myRealEscapeString($("#flagHash").val())
    validateFlagWthServer publicId, flagHash, (response) ->
      if response.msg == "flag validated"
        m.reload()
      alert response.msg

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

  # Setup a timer every 5 seconds
  setInterval model.reload, 5000

  model

$(document).ready(->
  server = new server()
  ko.applyBindings new scoreboardViewModel()
)
