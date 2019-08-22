import React, { Component } from 'react'
import { AudioState } from './components/utils'
import AudioPlayerSmall from './audio_player_small'
import AudioPlayerMedium from './audio_player_medium'
import AudioPlayerLarge from './audio_player_large'

interface Props {
  podcast: string
  episode: string
  podcastId: string
  episodeId: string
  albumArt: string
  audioSrc: string
  audioType: string
}

interface State {
  screen: 'none' | 'sm' | 'md' | 'lg'
  duration: number
  currentTime: number
  audioState: AudioState
  expandOnMobile: boolean
}

export default class AudioPlayer extends Component<Props, State> {
  state = {
    screen: 'none' as 'none' | 'sm' | 'md' | 'lg',
    duration: 0,
    currentTime: 0,
    audioState: 'LOADING' as AudioState,
    expandOnMobile: true,
  }

  audio: HTMLAudioElement = undefined

  componentDidMount() {
    this.audio = document.createElement('audio')

    // Play
    this.audio.addEventListener('canplay', () => {
      this.setState({ audioState: 'PAUSED' })
      this.audio.play()
    })
    this.audio.addEventListener('playing', () => {
      this.setState({ audioState: 'PLAYING' })
    })
    // Pause
    this.audio.addEventListener('pause', () => {
      this.setState({ audioState: 'PAUSED' })
    })
    // Loading
    this.audio.addEventListener('loadstart', () => {
      this.setState({ audioState: 'LOADING' })
    })
    this.audio.addEventListener('seeking', () => {
      this.setState({ audioState: 'LOADING' })
    })
    //Ended
    this.audio.addEventListener('ended', () => {
      this.setState({ audioState: 'ENDED' })
    })
    // Duration
    this.audio.addEventListener('durationchange', () => {
      this.setState({ duration: this.audio.duration })
    })
    // Current time
    this.audio.addEventListener('timeupdate', () => {
      this.setState({ currentTime: this.audio.currentTime })
    })

    this.audio.src = this.props.audioSrc

    // window resize
    window.addEventListener('resize', () => {
      this.handleScreenResize()
    })

    this.handleScreenResize()
  }

  componentWillUnmount() {
    window.removeEventListener('resize', () => {})
  }

  handleScreenResize = () => {
    const screenWidth = window.innerWidth
    if (screenWidth >= 1024) {
      this.setState({ screen: 'lg' })
      return
    }
    if (screenWidth >= 768) {
      this.setState({ screen: 'md' })
      return
    }
    this.setState({ screen: 'sm' })
  }

  handleActionButtonPress = () => {
    const { audioState } = this.state
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
    const { screen, audioState, currentTime, duration } = this.state

    const { podcast, episode, podcastId, episodeId, albumArt } = this.props

    if (screen === 'sm') {
      return (
        <AudioPlayerSmall
          podcast={podcast}
          episode={episode}
          podcastId={podcastId}
          episodeId={episodeId}
          albumArt={albumArt}
          audioState={audioState}
          duration={duration}
          currentTime={currentTime}
          handleSeek={this.handleSeek}
          handleFastForward={this.handleFastForward}
          handleActionButtonPress={this.handleActionButtonPress}
        />
      )
    }

    if (screen === 'md') {
      return (
        <AudioPlayerMedium
          podcast={podcast}
          episode={episode}
          podcastId={podcastId}
          episodeId={episodeId}
          albumArt={albumArt}
          audioState={audioState}
          duration={duration}
          currentTime={currentTime}
          handleSeek={this.handleSeek}
          handleFastForward={this.handleFastForward}
          handleActionButtonPress={this.handleActionButtonPress}
        />
      )
    }

    if (screen === 'lg') {
      return (
        <AudioPlayerLarge
          podcast={podcast}
          episode={episode}
          podcastId={podcastId}
          episodeId={episodeId}
          albumArt={albumArt}
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
