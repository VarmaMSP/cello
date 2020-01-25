import ButtonShowMore from 'components/button_show_more'
import EpisodePreview from 'components/episode_preview'
import React, { useEffect } from 'react'
import { Episode, Podcast } from 'types/app'

export interface StateToProps {
  isUserSignedIn: boolean
  podcast: Podcast
  episodes: Episode[]
  receivedAll: boolean
  isLoadingMore: boolean
}

export interface DispatchToProps {
  loadPlaybacks: (episodeIds: string[]) => void
  loadEpisodes: (offset: number) => void
}

export interface OwnProps {
  podcastId: string
}

interface Props extends StateToProps, DispatchToProps, OwnProps {}

const ListEpisodes: React.SFC<Props> = ({
  isUserSignedIn,
  podcast,
  episodes,
  loadPlaybacks,
  loadEpisodes,
  receivedAll,
  isLoadingMore,
}) => {
  useEffect(() => {
    loadPlaybacks(episodes.map((e) => e.id))
  }, [])

  useEffect(() => {
    loadPlaybacks(episodes.map((e) => e.id))
  }, [isUserSignedIn])

  return (
    <>
      {episodes.map((episode) => (
        <div key={episode.id} className="mb-4">
          <EpisodePreview episodeId={episode.id} small showIcon />
        </div>
      ))}
      {episodes.length < podcast.totalEpisodes && !receivedAll && (
        <div className="w-full h-10 mx-auto my-6">
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
