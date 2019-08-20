import React, { Component } from 'react'
import Controls from './components/controls'
import ProgressBar from './components/progress_bar'
import { threadId } from 'worker_threads';

interface Props {
  audioSrc: string;
  audioType: string;
  playWhenReady: boolean;
}

interface State {
  state: "PAUSED" | "PLAYING" | "LOADING" | "ENDED";
  duration: number;
  currentTime: number;
}

export default class PlayerX extends Component<Props, State> {
  state = {
    state: "LOADING" as "PAUSED" | "PLAYING" | "LOADING" | "ENDED",
    duration: 0,
    currentTime: 0,
  }

  audio: HTMLAudioElement = undefined
  
  componentDidMount() {
    this.audio = document.createElement('audio')

    // Play
    this.audio.addEventListener('canplay', () => {
      this.audio.play()
      this.setState({state: 'PLAYING'})
    })
    this.audio.addEventListener('canplaythrough', () => {
      this.audio.play()
      this.setState({state: 'PLAYING'})
    }) 
    this.audio.addEventListener('play', () =>
      this.setState({state: 'PLAYING'})
    )
    this.audio.addEventListener('playing', () =>
      this.setState({state: 'PLAYING'})
    )

    // Pause
    this.audio.addEventListener('pause', () =>
      this.setState({state: 'PAUSED'})
    )

    // Loading
    this.audio.addEventListener('loadstart', () => 
      this.setState({state: 'LOADING'})
    )
    this.audio.addEventListener('seeking', () => 
      this.setState({state: 'LOADING'})
    )

    //Ended
    this.audio.addEventListener("ended", () => {
      this.setState({state: "ENDED"})
    })

    // Duration
    this.audio.addEventListener("durationchange", () => {
      this.setState({duration: this.audio.duration})
    })

    // Current time
    this.audio.addEventListener('timeupdate', () =>
      this.setState({currentTime: this.audio.currentTime})
    )

    this.audio.src = "http://traffic.libsyn.com/joeroganexp/p1335.mp3?dest-id=19997"
  }

  handlePlayToggle = () => {
    const { state } = this.state
    if (state === "PLAYING") {
      this.audio.pause()
    }
    if (state === "PAUSED") {
      this.audio.play()
    }
  }

  handleSeek = (t: number) => {
    this.audio.currentTime = t
    this.setState({
      state: "LOADING",
      currentTime: t,
    })
  }

  handleFastForward = (t: number) => {
    const { currentTime } = this.state

    this.audio.currentTime = currentTime + t
    this.setState({
      state: "LOADING",
      currentTime: currentTime + t,
    })
  }
  
  render() {
    const {currentTime, duration} = this.state

    return <div className="w-full px-8 mt-64">
      <Controls 
        playing={this.state.state === "PLAYING"}
        handlePlayToggle={this.handlePlayToggle}
        handleFastForward={this.handleFastForward}/>
      <ProgressBar
        currentTime={currentTime}
        duration={duration}
        handleSeek={this.handleSeek}
      />
    </div>
  }
}