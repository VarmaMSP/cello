import React from 'react'
import Shiitake from 'shiitake'
import ButtonWithIcon from '../button_with_icon'
import { Episode } from '../../types/app'
import * as Utils from '../utils'

export interface StateToProps {
  episodes: Episode[]
}

export interface DispatchToProps {
  playEpisode: (episodeId: string) => void
}

export interface OwnProps {
  podcastId: string
}

interface Props extends StateToProps, DispatchToProps, OwnProps {}

const EpisodeList: React.SFC<Props> = ({ episodes, playEpisode }) => {
  return (
    <>
      {episodes.map(({ id, title, duration, pubDate }) => (
        <div key={id}>
          <div className="flex justify-between my-2 lg:px-6 sm:px-4 py-2 rounded-full lg:hover:bg-gray-200">
            <div className="flex-auto">
              <Shiitake
                lines={2}
                throttleRate={200}
                className="md:text-base text-sm tracking-wide leading-relaxed"
              >
                {title}
              </Shiitake>
              <span className="text-xs">
                {Utils.humanizeDuration(duration)}
                <span className="mx-2 font-extrabold">&middot;</span>
                {Utils.humanizePastDate(pubDate)}
              </span>
            </div>
            <ButtonWithIcon
              className="flex-none w-6 ml-2 text-gray-700"
              icon="play-outline"
              onClick={() => playEpisode(id)}
            />
          </div>
        </div>
      ))}
    </>
  )
}

export default EpisodeList
