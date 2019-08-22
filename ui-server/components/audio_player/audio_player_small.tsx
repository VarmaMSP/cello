import React, { Component } from 'react'
import { AudioState } from './components/utils'
import ButtonWithIcon from '../button_with_icon'
import ActionButton from './components/action_button'
import NavbarBottom from '../navbar_bottom/navbar_bottom'
import ProgressBar from './components/progress_bar'

interface Props {
  podcast: string
  episode: string
  podcastId: string
  episodeId: string
  albumArt: string
  audioState: AudioState
  duration: number
  currentTime: number
  handleSeek: (t: number) => void
  handleFastForward: (t: number) => void
  handleActionButtonPress: () => void
}

interface State {
  expand: boolean
}

export default class AudioPlayerSmall extends Component<Props, State> {
  state = {
    expand: false,
  }

  handleExpandToggle = () => {
    const { expand } = this.state
    this.setState({ expand: !expand })
  }

  render() {
    const { expand } = this.state
    const {
      podcast,
      episode,
      albumArt,
      audioState,
      currentTime,
      duration,
      handleActionButtonPress,
      handleSeek,
      handleFastForward,
    } = this.props

    return (
      <footer
        className="fixed bottom-0 left-0 flex flex-col justify-between bg-white w-full h-auto border-t"
        style={{
          willChange: 'height',
          height: expand ? '24rem' : '8rem',
          transition: '0.4s cubic-bezier(.22,.86,.62,.95)',
        }}
      >
        {/* Player minimised */}
        <section className="player justify-between flex h-16 w-full px-2">
          <div
            className="details flex-auto px-2 pt-3 mr-1 truncate"
            style={{
              opacity: expand ? 0 : 1,
              transition: '0.1s',
            }}
          >
            <h4 className="text-sm font-semibold leading-relaxed">{episode}</h4>
            <h4 className="text-xs leading-relaxed">{podcast}</h4>
          </div>
          <div className="flex-none flex items-center">
            <div
              className="-mr-3"
              style={{
                display: expand ? 'none' : 'block',
                transition: '0.1s',
              }}
            >
              <ActionButton
                audioState={audioState}
                handleActionButtonPress={handleActionButtonPress}
              />
            </div>
            <div
              style={{
                transform: expand ? 'rotate(180deg)' : 'none',
                transition: '0.4s ease-out',
              }}
            >
              <ButtonWithIcon
                className="w-10 mx-1"
                icon="cheveron-up"
                onClick={this.handleExpandToggle}
              />
            </div>
          </div>
        </section>

        {/* Player Maximised */}
        <section
          className="player-max flex flex-col justify-start w-full mb-6"
          style={{
            willChange: 'height',
            height: expand ? '100%' : '0%',
            marginBottom: expand ? '2rem' : '0rem',
            transition: '0.4s cubic-bezier(.22,.86,.62,.95)',
          }}
        >
          <section className="px-3 flex flex-row">
            <img
              className="h-32 w-32 flex-none object-cover object-center rounded"
              src={albumArt}
            />
            <section className="flex-1 ml-3">
              <h4 className="text-md font-semibold text-green-600 leading-loose">
                Now Playing
              </h4>
              <h5 className="text-md font-bold text-gray-700 leading-tight">
                {episode}
              </h5>
              <h5 className="text-sm font-semibold text-gray-600 leading-loose">
                {podcast}
              </h5>
            </section>
          </section>
          <section className="flex flex-row items-center justify-center my-2 text-gray-800">
            <ButtonWithIcon
              className="w-10"
              icon="fast-rewind"
              onClick={() => handleFastForward(-10)}
            />
            <ActionButton
              audioState={audioState}
              handleActionButtonPress={handleActionButtonPress}
            />
            <ButtonWithIcon
              className="w-10"
              icon="fast-forward"
              onClick={() => handleFastForward(10)}
            />
          </section>
          <section className="px-4">
            <ProgressBar
              currentTime={currentTime}
              duration={duration}
              handleSeek={handleSeek}
            />
          </section>
        </section>

        {/* Bottom navbar */}
        <div
          className="overflow-y-hidden"
          style={{
            willChange: 'height',
            height: expand ? '0rem' : '4rem',
            opacity: expand ? 0 : 1,
            transition: '0.4s cubic-bezier(.22,.86,.12,1.01)',
          }}
        >
          <NavbarBottom />
        </div>
      </footer>
    )
  }
}
