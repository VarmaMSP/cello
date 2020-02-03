import ButtonWithIcon from 'components/button_with_icon'
import { EpisodeLink } from 'components/link'
import React, { Component } from 'react'
import { AudioState } from 'types/app'
import { Episode, Podcast } from 'types/models'
import { getImageUrl } from 'utils/dom'
import ActionButton from './components/action_button'
import NavbarBottom from './components/navbar_bottom'
import SeekBar from './components/seek_bar'

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

export default class AudioPlayerSmall extends Component<Props> {
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
          <h4 className="text-sm font-semibold text-gray-200 leading-relaxed">
            {episode.title}
          </h4>
          <h4 className="text-xs text-gray-300 leading-relaxed">
            {podcast.title}
          </h4>
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
              theme="dark"
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
              className="w-10 mx-1 text-gray-300"
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
        <section className="flex flex-row px-3 mb-2">
          <img
            className="h-32 w-32 flex-none object-cover object-center rounded border  "
            src={getImageUrl(podcast.urlParam)}
          />
          <EpisodeLink episodeUrlParam={episode.urlParam}>
            <a className="block flex-1 ml-3">
              <h5 className="text-sm font-bold text-gray-200 line-clamp-2">
                {episode.title}
              </h5>
              <h5 className="text-xs font-semibold text-gray-300 line-clamp-2">
                {podcast.title}
              </h5>
            </a>
          </EpisodeLink>
        </section>
        <section className="flex flex-row items-center justify-center my-2 text-gray-800">
          <ButtonWithIcon
            className="w-8 text-gray-300"
            icon="fast-rewind"
            onClick={() => handleFastForward(-10)}
          />
          <ActionButton
            theme="dark"
            audioState={audioState}
            handleActionButtonPress={handleActionButtonPress}
          />
          <ButtonWithIcon
            className="w-8 text-gray-300"
            icon="fast-forward"
            onClick={() => handleFastForward(10)}
          />
        </section>
        <section className="px-4">
          <SeekBar
            theme="dark"
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
        className="fixed bottom-0 left-0 flex flex-col justify-between w-full border-t border-gray-300 bg-black"
        style={{
          willChange: 'height',
          height: expandOnMobile ? '24rem' : episode ? '7.3rem' : '3.3rem',
          transition: '0.4s cubic-bezier(.22,.86,.62,.95)',
        }}
      >
        {playerMin}
        {playerMax}

        {/* Bottom navbar */}
        <div
          className="overflow-y-hidden z-40"
          style={{
            willChange: 'height',
            height: expandOnMobile ? '0rem' : '3.3rem',
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
