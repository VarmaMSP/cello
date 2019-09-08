import React, { Component } from 'react'
import { AudioState, Episode, Podcast, ScreenWidth } from 'types/app'
import AudioPlayerLarge from './audio_player_large'
import AudioPlayerMedium from './audio_player_medium'
import AudioPlayerSmall from './audio_player_small'

export interface StateToProps {
  episodeId: string
  episode: Episode
  podcast: Podcast
  audioState: AudioState
  screenWidth: ScreenWidth
  expandOnMobile: boolean
}

export interface DispatchToProps {
  setAudioState: (s: AudioState) => void
  toggleExpandOnMobile: () => void
}

interface State {
  duration: number
  currentTime: number
}

export default class AudioPlayer extends Component<
  StateToProps & DispatchToProps,
  State
> {
  state = {
    duration: 0,
    currentTime: 0,
  }

  // Avoid creating audio element in the constructor
  // Because document obect is not available on the server
  audio: HTMLAudioElement = {} as any

  componentDidUpdate(prevProps: StateToProps & DispatchToProps) {
    const { episodeId } = this.props
    if (episodeId !== '' && episodeId !== prevProps.episodeId) {
      this.audio.src = this.props.episode.mediaUrl
    }
  }

  componentDidMount() {
    this.audio = document.createElement('audio')

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
      this.setState({ duration: this.audio.duration })
    })
    // Current time
    this.audio.addEventListener('timeupdate', () => {
      this.setState({ currentTime: this.audio.currentTime })
    })

    if (this.props.episodeId !== '') {
      this.audio.src = this.props.episode.mediaUrl
    }
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
    this.setState({ currentTime: t })
  }

  handleFastForward = (t: number) => {
    const { currentTime } = this.state
    this.audio.currentTime = currentTime + t
    this.setState({ currentTime: currentTime + t })
  }

  render() {
    const { currentTime, duration } = this.state
    const {
      podcast,
      episode,
      audioState,
      screenWidth,
      expandOnMobile,
    } = this.props

    if (screenWidth === 'SM') {
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

    if (screenWidth === 'MD') {
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

    if (screenWidth === 'LG') {
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
