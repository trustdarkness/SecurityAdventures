myRealEscapeString = (str) ->
  return str.replace /[\0\x08\x09\x1a\n\r"'\\\%]/g, (char) ->
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

server = ->
  
  url = (route) -> "http://localhost:80/#{route}"
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
    model.sortUsers()

reloadModel = (model) ->
  server.fetch "scoreboard", ({scoreboard: data}) ->
    loaded = false
    for score in data.scores
      userMatch = ko.utils.arrayFirst model.users(), (item) ->
        score.public_user.public_id == item.id
      if userMatch?
        if userMatch.flags().length != score.public_flags.length
          userMatch.flags score.public_flags
          loaded = true
    if loaded then model.sortUsers()

userModel = (d) ->
  m =
    id: d.public_user.public_id
    flags: ko.observableArray d.public_flags

  m.flagsFound = ko.computed ->
    m.flags().length

  m.score = ko.computed ->
    score = 0
    for flag in m.flags()
      score += flag.value
    score

  m

sortUsers = (model) ->
  model.users.sort (left, right) ->
    if left.score() < right.score()
      return 1
    if right.score() < left.score()
      return -1
    else return 0

scoreboardModel = ->
  validateFlagWthServer = (publicId, flagHash, cb) ->
    validateFlagModel = { public_user_id: publicId, tag: flagHash }
    server.update "validateFlag", cb, validateFlagModel

  m =
    users: ko.observableArray []

  m.load = -> loadModel m

  m.reload = -> reloadModel m

  m.sortUsers = -> sortUsers m

  m.validateFlag = (d, e) ->
    publicId = myRealEscapeString($("#publicId").val()) || ""
    flagHash = myRealEscapeString($("#flagHash").val()) || ""
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

padZero = (str) ->
  if str.length == 1
    return "0" + str
  else return str

startCountdown = ->
  _second = 1000;
  _minute = _second * 60;
  _hour = _minute * 60;
  _day = _hour * 24;
  
  now = new Date()
  andThen = new Date("Nov 25 2013 20:00:00 CST")
  diff = new Date(andThen - now)

  hours = Math.floor((diff % _day) / _hour);
  minutes = Math.floor((diff % _hour) / _minute);
  seconds = Math.floor((diff % _minute) / _second);

  hours = padZero(hours + "")
  minutes = padZero(minutes + "")
  seconds = padZero(seconds + "")
  $("#Counter").countdown({
    image: "../../images/digits.png",
    startTime: "#{hours}:#{minutes}:#{seconds}"
  });

scoreboardViewModel = () ->
  model = scoreboardModel()
  model.load()

  startCountdown()

  # Setup a timer every 5 seconds
  setInterval model.reload, 5000

  model

$(document).ready(->
  server = new server()
  ko.applyBindings new scoreboardViewModel()
)
