import React, { Component } from 'react'
import Utils, { TouchOrMouseEvent } from './utils'

interface Props {
  duration: number;
  currentTime: number;
  handleSeek: (t: number) => void;
}

interface State {
  seeking: boolean;
  firstRender: boolean;
  sliderPosition: number;
}

export default class ProgressBar extends Component<Props, State> {
  state = {
    seeking: false,
    firstRender: true,
    sliderPosition: 0
  }

  seekBarRef = React.createRef<HTMLDivElement>()

  componentDidMount() {
    document.addEventListener("mousemove", (e: MouseEvent) => {
      if (!this.state.seeking) return
      this.handleSeek(e as any)
    })

    document.addEventListener("mouseup", (e: MouseEvent) => {
      if (!this.state.seeking) return
      this.handleSeekComplete(e as any)
    })

    this.setState({firstRender: false})
  }

  getSeekBarPosition = (): {clientX: number, width: number} => {
    const rect = this.seekBarRef.current.getBoundingClientRect()
    return {
      clientX: rect.left,
      width: rect.width,
    }
  }

  getSliderPosition = (): number => {
    const {currentTime, duration} = this.props
    const {firstRender, seeking, sliderPosition} = this.state

    if (firstRender) {
      return 0
    }
    if (seeking) {
      return sliderPosition
    }
    return (currentTime / duration) * this.getSeekBarPosition().width
  }

  getProgressDetails = (): [string, string] => {
    const {currentTime, duration} = this.props
    const {firstRender, seeking, sliderPosition} = this.state

    if (firstRender || !seeking) {
      return Utils.formatTimeForDisplay(currentTime, duration)
    }
    return Utils.formatTimeForDisplay(
      (sliderPosition / this.getSeekBarPosition().width) * duration,
      duration
    )
  }

  getNewSliderPosition = (e: TouchOrMouseEvent): number => {
    const { clientX: clickX } = Utils.getClickPosition(e)
    const { clientX: seekBarX, width } = this.getSeekBarPosition()
    if (seekBarX <= clickX && clickX <= seekBarX + width) {
      return clickX - seekBarX
    }
    return this.state.sliderPosition
  }

  handleSeekBegin = (e: TouchOrMouseEvent) => {
    this.setState({
      seeking: true,
      sliderPosition: this.getNewSliderPosition(e),
    })
  }

  handleSeekComplete = (_: TouchOrMouseEvent) => {
    const { duration } = this.props
    const { sliderPosition } = this.state

    this.setState({
      seeking: false,
    })

    this.props.handleSeek(
      (sliderPosition / this.getSeekBarPosition().width) * duration
    )
  }

  handleSeek = (e: TouchOrMouseEvent) => {
    if (!this.state.seeking) return
    this.setState({
      sliderPosition: this.getNewSliderPosition(e)
    })
  }

  render() {
    const sliderPosition = this.getSliderPosition()
    const [t, T] = this.getProgressDetails()

    return <>
      <div className="flex justify-between items-center w-full px-2">
        <span className="text-sm text-gray-800 leading-relaxed tracking-wider select-none">{t}</span>
        <span className="text-sm text-gray-800 leading-relaxed tracking-wider select-none">{T}</span>
      </div>
      <div className="relative flex items-center w-full h-4 cursor-pointer select-none"
        onTouchStart={this.handleSeekBegin}
        onTouchMove={this.handleSeek}
        onTouchEnd={this.handleSeekComplete}
        onMouseDown={this.handleSeekBegin}
      >
        <div className="absolute left-0 w-full h-1 bg-gray-300 rounded select-none"
          ref={this.seekBarRef}
        />
        <div className="absolute left-0 w-10 h-1 bg-green-500 rounded select-none"
          style={{'transition': 'ease', 'width': `${sliderPosition}px`}}
        />
        <div className="absolute w-4 h-4 -ml-2 rounded-full bg-white border shadow-md z-50 select-none"
          style={{'transition': 'ease', 'left': `${sliderPosition}px`}}
        />
      </div>
    </>
  }
}