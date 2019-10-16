import React, { Component } from 'react'
import { AudioState, Episode, Podcast, ViewportSize } from 'types/app'
import AudioPlayerLarge from './audio_player_large'
import AudioPlayerMedium from './audio_player_medium'
import AudioPlayerSmall from './audio_player_small'

export interface StateToProps {
  episodeId: string
  episode: Episode
  podcast: Podcast
  duration: number
  audioState: AudioState
  currentTime: number
  viewportSize: ViewportSize
  expandOnMobile: boolean
}

export interface DispatchToProps {
  syncPlayback: (episodeId: string, currentTime: number) => void
  setDuration: (t: number) => void
  setAudioState: (s: AudioState) => void
  setCurrentTime: (t: number) => void
  toggleExpandOnMobile: () => void
}

export default class AudioPlayer extends Component<
  StateToProps & DispatchToProps
> {
  // Avoid creating audio element in the constructor
  // Because document obect is not available on the server
  audio: HTMLAudioElement = {} as any

  componentDidUpdate(prevProps: StateToProps & DispatchToProps) {
    const { episodeId } = this.props
    if (episodeId !== '' && episodeId !== prevProps.episodeId) {
      this.audio.src = this.props.episode.mediaUrl
      this.audio.currentTime = this.props.currentTime
    }
  }

  componentDidMount() {
    this.audio = document.createElement('audio')
    this.audio.preload = 'auto'

    // Play
    this.audio.addEventListener('canplay', () => {
      this.props.setAudioState('PAUSED')
      this.audio.play()
    })
    this.audio.addEventListener('playing', () => {
      this.props.setAudioState('PLAYING')
    })
    // Pause
    this.audio.addEventListener('pause', () => {
      this.props.setAudioState('PAUSED')
    })
    // Loading
    this.audio.addEventListener('loadstart', () => {
      this.props.setAudioState('LOADING')
    })
    this.audio.addEventListener('seeking', () => {
      this.props.setAudioState('LOADING')
    })
    //Ended
    this.audio.addEventListener('ended', () => {
      this.props.setAudioState('ENDED')
    })
    // Duration
    this.audio.addEventListener('durationchange', () => {
      this.props.setDuration(this.audio.duration)
    })
    // Current time
    this.audio.addEventListener('timeupdate', () => {
      if (
        Math.floor(this.audio.currentTime) >=
        Math.floor(this.props.currentTime) + 1
      ) {
        this.props.setCurrentTime(this.audio.currentTime)
      }
    })

    if (this.props.episodeId !== '') {
      this.audio.src = this.props.episode.mediaUrl
    }

    setInterval(() => {
      const { episodeId, syncPlayback, audioState, currentTime } = this.props
      if (audioState === 'PLAYING') {
        syncPlayback(episodeId, currentTime)
      }
    }, 5000)
  }

  handleActionButtonPress = () => {
    const { audioState } = this.props
    if (audioState === 'PLAYING') {
      this.audio.pause()
    }
    if (audioState === 'PAUSED') {
      this.audio.play()
    }
    if (audioState === 'LOADING') {
      return
    }
  }

  handleSeek = (t: number) => {
    this.audio.currentTime = t
    this.props.setCurrentTime(t)
  }

  handleFastForward = (t: number) => {
    const { currentTime, setCurrentTime } = this.props
    this.audio.currentTime = currentTime + t
    setCurrentTime(currentTime)
  }

  render() {
    const {
      podcast,
      episode,
      duration,
      audioState,
      currentTime,
      viewportSize,
      expandOnMobile,
    } = this.props

    if (viewportSize === 'SM') {
      return (
        <AudioPlayerSmall
          podcast={podcast}
          episode={episode}
          audioState={audioState}
          duration={duration}
          currentTime={currentTime}
          expandOnMobile={expandOnMobile}
          handleSeek={this.handleSeek}
          handleFastForward={this.handleFastForward}
          handleActionButtonPress={this.handleActionButtonPress}
          toggleExpandOnMobile={this.props.toggleExpandOnMobile}
        />
      )
    }

    if (viewportSize === 'MD') {
      return (
        <AudioPlayerMedium
          podcast={podcast}
          episode={episode}
          audioState={audioState}
          duration={duration}
          currentTime={currentTime}
          handleSeek={this.handleSeek}
          handleFastForward={this.handleFastForward}
          handleActionButtonPress={this.handleActionButtonPress}
        />
      )
    }

    if (viewportSize === 'LG') {
      return (
        <AudioPlayerLarge
          podcast={podcast}
          episode={episode}
          audioState={audioState}
          duration={duration}
          currentTime={currentTime}
          handleSeek={this.handleSeek}
          handleFastForward={this.handleFastForward}
          handleActionButtonPress={this.handleActionButtonPress}
        />
      )
    }

    return <></>
  }
}
