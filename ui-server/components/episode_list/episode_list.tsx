import ButtonWithIcon from 'components/button_with_icon'
import React from 'react'
import { Episode } from 'types/app'
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
        <div
          key={id}
          className="flex justify-between mb-2 lg:pl-6 lg:pr-3 py-2 rounded-full lg:hover:bg-gray-200"
        >
          <div className="w-11/12 pr-3">
            <span className="text-xs text-gray-600">
              {Utils.humanizePastDate(pubDate)}
              <span className="mx-2 font-extrabold">&middot;</span>
              {Utils.humanizeDuration(duration)}
            </span>
            <p className="font-medium md:text-base text-sm text-gray-800 tracking-wide truncate">
              {title}
            </p>
          </div>
          <ButtonWithIcon
            className="w-8 mx-auto text-gray-600 hover:text-black"
            icon="play-outline"
            onClick={() => playEpisode(id)}
          />
        </div>
      ))}
    </>
  )
}

export default EpisodeList
