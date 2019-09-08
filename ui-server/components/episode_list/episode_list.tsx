import React from 'react'
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
          <div className="flex justify-between my-2 mx-1 rounded-full hover:bg-gray-100">
            <div className="flex-auto">
              <h4 className="text-sm font-medium">{title}</h4>
              <span className="text-xs">
                {Utils.humanizeDuration(duration)}
                <span className="mx-2 font-extrabold">&middot;</span>
                {Utils.humanizePastDate(pubDate)}
              </span>
            </div>
            <ButtonWithIcon
              className="w-5"
              icon="play"
              onClick={() => playEpisode(id)}
            />
          </div>
          <hr className="my-3" />
        </div>
      ))}
    </>
  )
}

export default EpisodeList
