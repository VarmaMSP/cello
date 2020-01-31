import React, { Component } from 'react'
import { AudioState, Episode, Podcast, ViewportSize } from 'types/app'
import AudioPlayerLarge from './audio_player_large'
import AudioPlayerSmall from './audio_player_small'

export interface StateToProps {
  episodeId: string
  episode: Episode
  podcast: Podcast
  duration: number
  audioState: AudioState
  currentTime: number
  volume: number
  playbackRate: number
  viewportSize: ViewportSize
  expandOnMobile: boolean
}

export interface DispatchToProps {
  syncPlayback: (episodeId: string, position: number) => void
  setDuration: (t: number) => void
  setAudioState: (s: AudioState) => void
  setCurrentTime: (t: number) => void
  setVolume: (v: number) => void
  setPlaybackRate: (p: number) => void
  toggleExpandOnMobile: () => void
}

interface Props extends StateToProps, DispatchToProps {}

export default class AudioPlayer extends Component<Props> {
  // Avoid creating audio element in the constructor
  // Because document obect is not available on the server
  audio: HTMLAudioElement = {} as any

  componentDidUpdate(prevProps: Props) {
    const { episodeId, volume, playbackRate } = this.props

    if (episodeId !== '' && episodeId !== prevProps.episodeId) {
      this.audio.src = this.props.episode.mediaUrl
      this.audio.currentTime = this.props.currentTime
      this.props.syncPlayback(
        prevProps.episodeId,
        (prevProps.currentTime / prevProps.duration) * 100,
      )
    }

    if (Math.abs(volume - prevProps.volume) < 1e-5 ) {
      this.audio.volume = volume
    }

    if (Math.abs(playbackRate - prevProps.playbackRate) < 1e-5) {
      this.audio.playbackRate = playbackRate
    }
  }

  componentDidMount() {
    this.audio = document.createElement('audio')
    this.audio.preload = 'auto'
    this.audio.autoplay = true
    this.audio.volume = this.props.volume
    this.audio.playbackRate = this.props.playbackRate

    // Can Play
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
      this.props.syncPlayback(
        this.props.episodeId,
        (this.audio.currentTime / this.audio.duration) * 100,
      )
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
      this.props.syncPlayback(this.props.episodeId, 100)
    })
    // Duration
    this.audio.addEventListener('durationchange', () => {
      this.props.setDuration(this.audio.duration)
    })
    // Current time
    this.audio.addEventListener('timeupdate', () => {
      this.props.setAudioState('PLAYING')
      this.props.setCurrentTime(this.audio.currentTime)
    })

    if (this.props.episodeId !== '') {
      this.audio.src = this.props.episode.mediaUrl
    }

    setInterval(() => {
      const {
        episodeId,
        syncPlayback,
        duration,
        audioState,
        currentTime,
      } = this.props
      if (audioState === 'PLAYING') {
        syncPlayback(episodeId, (currentTime / duration) * 100)
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

  handleVolumeChange = (v: number) => {
    this.audio.volume = v
    this.props.setVolume(v)
  }

  handlePlaybackRateChange = (r: number) => {
    this.audio.playbackRate = r
    this.props.setPlaybackRate(r)
  }

  render() {
    const {
      podcast,
      episode,
      duration,
      audioState,
      currentTime,
      volume,
      playbackRate,
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

    if (viewportSize === 'MD' || viewportSize === 'LG') {
      return (
        <AudioPlayerLarge
          podcast={podcast}
          episode={episode}
          audioState={audioState}
          duration={duration}
          currentTime={currentTime}
          volume={volume}
          playbackRate={playbackRate}
          handleSeek={this.handleSeek}
          handleFastForward={this.handleFastForward}
          handleVolumeChange={this.handleVolumeChange}
          handlePlaybackRateChange={this.handlePlaybackRateChange}
          handleActionButtonPress={this.handleActionButtonPress}
        />
      )
    }

    return <></>
  }
}
