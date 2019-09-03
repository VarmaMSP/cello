import React, { Component } from 'react'
import ButtonWithIcon from '../button_with_icon'
import ActionButton from './components/action_button'
import NavbarBottom from '../navbar_bottom/navbar_bottom'
import ProgressBar from './components/progress_bar'
import { AudioState, Podcast, Episode } from '../../types/app'

interface Props {
  episode: Episode
  podcast: Podcast
  audioState: AudioState
  duration: number
  currentTime: number
  expandOnMobile: boolean
  handleSeek: (t: number) => void
  handleFastForward: (t: number) => void
  handleActionButtonPress: () => void
  toggleExpandOnMobile: () => void
}

interface State {}

export default class AudioPlayerSmall extends Component<Props, State> {
  render() {
    const {
      podcast,
      episode,
      audioState,
      duration,
      currentTime,
      expandOnMobile,
      handleActionButtonPress,
      handleSeek,
      handleFastForward,
    } = this.props

    const playerMin = episode ? (
      <section className="player justify-between flex h-16 w-full px-2">
        <div
          className="details flex-auto px-2 pt-3 mr-1 truncate"
          style={{
            opacity: expandOnMobile ? 0 : 1,
            transition: '0.1s',
          }}
        >
          <h4 className="text-sm font-semibold leading-relaxed">
            {episode.title}
          </h4>
          <h4 className="text-xs leading-relaxed">{podcast.title}</h4>
        </div>
        <div className="flex-none flex items-center">
          <div
            className="-mr-3"
            style={{
              display: expandOnMobile ? 'none' : 'block',
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
              transform: expandOnMobile ? 'rotate(180deg)' : 'none',
              transition: '0.4s ease-out',
            }}
          >
            <ButtonWithIcon
              className="w-10 mx-1"
              icon="cheveron-up"
              onClick={this.props.toggleExpandOnMobile}
            />
          </div>
        </div>
      </section>
    ) : (
      <></>
    )

    const playerMax = episode ? (
      <section
        className="player-max flex flex-col justify-start w-full mb-6"
        style={{
          willChange: 'height',
          height: expandOnMobile ? '100%' : '0%',
          marginBottom: expandOnMobile ? '2rem' : '0rem',
          transition: '0.4s cubic-bezier(.22,.86,.62,.95)',
        }}
      >
        <section className="px-3 flex flex-row">
          <img className="h-32 w-32 flex-none object-cover object-center rounded" />
          <section className="flex-1 ml-3">
            <h4 className="text-md font-semibold text-green-600 leading-loose">
              Now Playing
            </h4>
            <h5 className="text-md font-bold text-gray-700 leading-tight">
              {episode.title}
            </h5>
            <h5 className="text-sm font-semibold text-gray-600 leading-loose">
              {podcast.title}
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
    ) : (
      <></>
    )

    return (
      <footer
        className="fixed bottom-0 left-0 flex flex-col justify-between bg-white w-full  border-t border-gray-400"
        style={{
          willChange: 'height',
          height: expandOnMobile ? '24rem' : episode ? '8rem' : '4rem',
          transition: '0.4s cubic-bezier(.22,.86,.62,.95)',
        }}
      >
        {/* Player minimised */}
        {playerMin}

        {/* Player Maximised */}
        {playerMax}

        {/* Bottom navbar */}
        <div
          className="overflow-y-hidden"
          style={{
            willChange: 'height',
            height: expandOnMobile ? '0rem' : '4rem',
            opacity: expandOnMobile ? 0 : 1,
            transition: '0.4s cubic-bezier(.22,.86,.12,1.01)',
          }}
        >
          <NavbarBottom />
        </div>
      </footer>
    )
  }
}
