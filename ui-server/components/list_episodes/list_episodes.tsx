import ButtonWithIcon from 'components/button_with_icon'
import React from 'react'
import { Episode } from 'types/app'
import { formatEpisodeDuration, formatEpisodePubDate } from 'utils/format'

export interface StateToProps {
  episodes: Episode[]
}

export interface DispatchToProps {
  playEpisode: (episodeId: string) => void
  showEpisodeModal: (episodeId: string) => void
}

export interface OwnProps {
  podcastId: string
}

interface Props extends StateToProps, DispatchToProps, OwnProps {}

const ListEpisodes: React.SFC<Props> = ({
  episodes,
  playEpisode,
  showEpisodeModal,
}) => {
  return (
    <>
      {episodes.map(({ id, title, duration, pubDate }) => (
        <div
          key={id}
          className="flex mb-3 lg:px-6 py-2 rounded-full lg:hover:bg-gray-200"
        >
          <div className="flex-auto w-11/12 pr-3">
            <span className="text-xs text-gray-700">
              {formatEpisodePubDate(pubDate)}
              <span className="mx-2 font-extrabold">&middot;</span>
              {formatEpisodeDuration(duration)}
            </span>
            <p className="pb-1 font-medium md:text-base text-sm text-gray-800 tracking-wide truncate">
              <span
                className="cursor-pointer hover:underline"
                onClick={() => showEpisodeModal(id)}
              >
                {title}
              </span>
            </p>
          </div>
          <ButtonWithIcon
            className="flex-none md:w-8 w-6 mx-auto text-gray-600 hover:text-black"
            icon="play-outline"
            onClick={() => playEpisode(id)}
          />
        </div>
      ))}
    </>
  )
}

export default ListEpisodes
