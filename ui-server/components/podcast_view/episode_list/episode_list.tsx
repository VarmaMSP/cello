import ButtonPlay from 'components/button_play'
import ButtonShowMore from 'components/button_show_more'
import ButtonWithIcon from 'components/button_with_icon'
import EpisodeMeta from 'components/episode_meta'
import React, { useEffect } from 'react'
import striptags from 'striptags'
import { Episode } from 'types/app'

export interface StateToProps {
  episodes: Episode[]
  receivedAllEpisodes: boolean
  isLoadingMore: boolean
}

export interface DispatchToProps {
  showAddToPlaylistModal: (episode: string) => void
  loadEpisodePlaybacks: (episodeIds: string[]) => void
  loadEpisodes: (offset: number) => void
}

export interface OwnProps {
  podcastId: string
}

interface Props extends StateToProps, DispatchToProps, OwnProps {}

const ListEpisodes: React.SFC<Props> = ({
  episodes,
  showAddToPlaylistModal,
  loadEpisodePlaybacks,
  loadEpisodes,
  receivedAllEpisodes,
  isLoadingMore,
}) => {
  useEffect(() => {
    loadEpisodePlaybacks(episodes.map((e) => e.id))
  }, [])

  return (
    <>
      {episodes.map((episode) => (
        <div key={episode.id}>
          <div className="flex group mb-3 py-2 cursor-default">
            <div className="flex-auto w-11/12">
              <EpisodeMeta episodeId={episode.id} />
              <p
                className="md:text-base text-sm text-black tracking-wide truncate"
                style={{ marginTop: '1px' }}
              >
                {episode.title}
              </p>
              <p
                className="mt-1 text-sm text-gray-700 line-clamp-2 tracking-wide"
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
                onClick={() => showAddToPlaylistModal(episode.id)}
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
      {!receivedAllEpisodes && (
        <div className="w-28 h-10 mx-auto my-6">
          <ButtonShowMore
            isLoading={isLoadingMore}
            loadMore={() =>
              episodes.length > 0 && loadEpisodes(episodes.length)
            }
          />
        </div>
      )}
    </>
  )
}

export default ListEpisodes
