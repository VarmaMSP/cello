const playIcon =
  '<svg xmlns="http://www.w3.org/2000/svg" height="24" viewBox="0 0 24 24" width="24"><path d="M8 5v14l11-7z"/><path d="M0 0h24v24H0z" fill="none"/></svg>'
const pauseIcon =
  '<svg xmlns="http://www.w3.org/2000/svg" height="24" viewBox="0 0 24 24" width="24"><path d="M6 19h4V5H6v14zm8-14v14h4V5h-4z"/><path d="M0 0h24v24H0z" fill="none"/></svg>'

const audio = document.createElement('audio')

const audioProgress = () => {
  if (audio.duration == 0) return 0
  return (audio.currentTime / audio.duration) * 100
}

const audioDurationFormatted = () => {
  const format = (t) =>
    new Date(0, 0, 0, 0, 0, t || 0).toTimeString().match(/(\d\d:\d\d:\d\d)/)[1]
  return `${format(audio.currentTime)} | ${format(audio.duration)}`
}

document.getElementById('pause-play').onclick = () => {
  audio.paused ? audio.play() : audio.pause()
}

document.getElementById('fast-rewind').onclick = () => {
  audio.currentTime = audio.currentTime > 10 ? audio.currentTime - 10 : 0
}

document.getElementById('fast-forward').onclick = () => {
  audio.currentTime =
    audio.currentTime + 10 < audio.duration
      ? audio.currentTime + 10
      : audio.duration
}

audio.addEventListener('canplay', () => {
  document.getElementById('pause-play').innerHTML = playIcon
})

audio.addEventListener('playing', () => {
  document.getElementById('pause-play').innerHTML = pauseIcon
})

audio.addEventListener('pause', () => {
  document.getElementById('pause-play').innerHTML = playIcon
})

audio.addEventListener('seeking', () => {})

audio.addEventListener('timeupdate', () => {
  document.getElementById('duration').innerText = audioDurationFormatted()
  document.getElementById('progress').style['width'] = `${audioProgress()}%`
})

audio.addEventListener('durationchange', () => {
  document.getElementById('duration').innerText = audioDurationFormatted()
  document.getElementById('progress').style['width'] = `${audioProgress()}%`
})

const data = document.getElementById('DATA')
const mediaUrl = (data ? JSON.parse(data.innerHTML) : {})['media_url']

audio.src = mediaUrl
