import ButtonWithIcon from 'components/button_with_icon'
import ProgressBar from 'components/progress_bar'
import React, { useEffect } from 'react'
import { Episode, EpisodePlayback } from 'types/app'
import { formatEpisodeDuration, formatEpisodePubDate } from 'utils/format'

export interface StateToProps {
  episodes: Episode[]
  episodePlaybacks: { [episodeId: string]: EpisodePlayback }
}

export interface DispatchToProps {
  playEpisode: (episodeId: string) => void
  showEpisodeModal: (episodeId: string) => void
  loadEpisodePlaybacks: (episodeIds: string[]) => void
}

export interface OwnProps {
  podcastId: string
}

interface Props extends StateToProps, DispatchToProps, OwnProps {}

const ListEpisodes: React.SFC<Props> = ({
  episodes,
  episodePlaybacks,
  playEpisode,
  showEpisodeModal,
  loadEpisodePlaybacks,
}) => {
  useEffect(() => {
    loadEpisodePlaybacks(episodes.map((e) => e.id))
  }, [])

  return (
    <>
      {episodes.map(({ id, title, duration, pubDate }) => (
        <div
          key={id}
          className="flex mb-3 lg:px-6 py-2 rounded-full lg:hover:bg-gray-200"
        >
          <div className="flex-auto w-11/12 pr-3">
            <div className="flex items-center justify-between">
              <span className="text-xs text-gray-700">
                {formatEpisodePubDate(pubDate)}
                <span className="mx-2 font-extrabold">&middot;</span>
                {formatEpisodeDuration(duration)}
              </span>
              {episodePlaybacks[id] && (
                <div className="lg:w-1/3 w-2/5 lg:mr-6" title="progress">
                  <ProgressBar
                    currentTime={episodePlaybacks[id].currentTime}
                    duration={duration}
                  />
                </div>
              )}
            </div>

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
