import ButtonPlay from 'components/button_play'
import ButtonWithIcon from 'components/button_with_icon'
import EpisodeMeta from 'components/episode_meta'
import React, { useEffect } from 'react'
import striptags from 'striptags'
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
        <div key={episode.id}>
          <div className="flex group mb-3 py-2 cursor-default">
            <div className="flex-auto w-11/12 pr-3">
              <EpisodeMeta episodeId={episode.id} />
              <p
                className="md:text-base text-sm text-black tracking-wide truncate"
                onClick={() => showEpisodeModal(episode.id)}
                style={{ marginTop: '1px' }}
              >
                {episode.title}
              </p>
              <p
                className="mt-1 text-sm text-gray-600 line-clamp-3"
                style={{ hyphens: 'auto' }}
              >
                {striptags(episode.description)}
              </p>
            </div>
            <div className="flex flex-col items-center justify-end ml-4">
              <ButtonPlay className="w-6" episodeId={episode.id} />
              <ButtonWithIcon
                className="group-hover:block text-gray-600 w-5 my-4"
                icon="playlist-add"
              />
              <ButtonWithIcon
                className="group-hover:block text-gray-600 w-4"
                icon="share"
              />
            </div>
          </div>
          <hr className="my-3" />
        </div>
      ))}
    </>
  )
}

export default ListEpisodes
