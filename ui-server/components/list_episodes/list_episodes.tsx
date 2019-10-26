import ButtonPlay from 'components/button_play'
import EpisodeMeta from 'components/episode_meta'
import React, { useEffect } from 'react'
import { Episode } from 'types/app'

export interface StateToProps {
  episodes: Episode[]
}

export interface DispatchToProps {
  showEpisodeModal: (episodeId: string) => void
  loadEpisodePlaybacks: (episodeIds: string[]) => void
}

export interface OwnProps {
  podcastId: string
}

interface Props extends StateToProps, DispatchToProps, OwnProps {}

const ListEpisodes: React.SFC<Props> = ({
  episodes,
  showEpisodeModal,
  loadEpisodePlaybacks,
}) => {
  useEffect(() => {
    loadEpisodePlaybacks(episodes.map((e) => e.id))
  }, [])

  return (
    <>
      {episodes.map((episode) => (
        <div
          key={episode.id}
          className="flex mb-3 lg:px-6 py-2 rounded-full lg:hover:bg-gray-200 cursor-default"
        >
          <div className="flex-auto w-11/12 pr-3">
            <EpisodeMeta episodeId={episode.id} />
            <p
              className="font-medium md:text-base text-sm text-gray-800 tracking-wide truncate"
              onClick={() => showEpisodeModal(episode.id)}
              style={{ marginTop: '1px' }}
            >
              {episode.title}
            </p>
          </div>
          <ButtonPlay className="md:w-8 w-6" episodeId={episode.id} />
        </div>
      ))}
    </>
  )
}

export default ListEpisodes
