import ButtonShowMore from 'components/button_show_more'
import EpisodeListItem from 'components/episode_list_item'
import PodcastPreview from 'components/podcast_preview'
import React from 'react'
import { SearchResultType, SearchSortBy } from 'types/search'

export interface StateToProps {
  searchBarText: string
  resultType: SearchResultType
  sortBy: SearchSortBy
  podcastIds: string[]
  episodeIds: string[]
  receivedAll: boolean
  isLoadingMore: boolean
}

export interface DispatchToProps {
  loadMore: (
    a: string,
    b: SearchResultType,
    c: SearchSortBy,
    d: number,
    e: number,
  ) => void
}

const SearchResultsList: React.FC<StateToProps & DispatchToProps> = ({
  searchBarText,
  resultType,
  sortBy,
  podcastIds,
  episodeIds,
  receivedAll,
  isLoadingMore,
  loadMore,
}) => {
  if (resultType === 'podcast') {
    return (
      <div>
        {podcastIds.map((id) => (
          <PodcastPreview key={id} podcastId={id} />
        ))}

        {!receivedAll && (
          <div className="w-full h-10 mx-auto my-6">
            <ButtonShowMore
              isLoading={isLoadingMore}
              loadMore={() =>
                loadMore(
                  searchBarText,
                  resultType,
                  sortBy,
                  podcastIds.length,
                  20,
                )
              }
            />
          </div>
        )}
      </div>
    )
  }

  if (resultType === 'episode') {
    return (
      <div>
        {episodeIds.map((id) => (
          <EpisodeListItem episodeId={id} key={id} />
        ))}

        {!receivedAll && (
          <div className="w-full h-10 mx-auto my-6">
            <ButtonShowMore
              isLoading={isLoadingMore}
              loadMore={() =>
                loadMore(
                  searchBarText,
                  resultType,
                  sortBy,
                  episodeIds.length,
                  20,
                )
              }
            />
          </div>
        )}
      </div>
    )
  }

  return <></>
}

export default SearchResultsList
