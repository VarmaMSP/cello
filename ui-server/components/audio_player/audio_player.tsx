import React, { Component } from 'react'
import Controls from './components/controls'
import ProgressBar from './components/progress_bar'
import { AudioState } from './components/utils'

interface Props {
  audioSrc: string;
  audioType: string;
  playWhenReady: boolean;
}

interface State {
  duration: number;
  currentTime: number;
  audioState: AudioState;
}

export default class AudioPlayer extends Component<Props, State> {
  state = {
    duration: 0,
    currentTime: 0,
    audioState: "LOADING" as AudioState,
  }

  audio: HTMLAudioElement = undefined
  
  componentDidMount() {
    this.audio = document.createElement('audio')

    // Play
    this.audio.addEventListener('canplay', () => {
      this.setState({audioState: 'PAUSED'})
      this.audio.play()
    })
    this.audio.addEventListener('playing', () => {
      this.setState({audioState: 'PLAYING'})
    })
    // Pause
    this.audio.addEventListener('pause', () => {
      this.setState({audioState: 'PAUSED'})
    })
    // Loading
    this.audio.addEventListener('loadstart', () => {
      this.setState({audioState: 'LOADING'})
    })
    this.audio.addEventListener('seeking', () => {
      this.setState({audioState: 'LOADING'})
    })
    //Ended
    this.audio.addEventListener("ended", () => {
      this.setState({audioState: "ENDED"})
    })
    // Duration
    this.audio.addEventListener("durationchange", () => {
      this.setState({duration: this.audio.duration})
    })
    // Current time
    this.audio.addEventListener('timeupdate', () => {
      this.setState({currentTime: this.audio.currentTime})
    })

    this.audio.src = this.props.audioSrc
  }

  handleActionButtonPress = () => {
    const { audioState } = this.state
    if (audioState === "PLAYING") {
      this.audio.pause()
    }
    if (audioState === "PAUSED") {
      this.audio.play()
    }
    if (audioState === "LOADING") {
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
    this.setState({currentTime: currentTime + t})
  }
  
  render() {
    const {currentTime, duration} = this.state

    return <div className="w-full px-8 mt-64">
      <Controls 
        audioState={this.state.audioState}
        handleFastForward={this.handleFastForward}
        handleActionButtonPress={this.handleActionButtonPress}/>
      <ProgressBar
        currentTime={currentTime}
        duration={duration}
        handleSeek={this.handleSeek}
      />
    </div>
  }
}