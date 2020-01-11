import ButtonShowMore from 'components/button_show_more'
import EpisodeListItem from 'components/episode_list_item'
import PodcastPreview from 'components/podcast_preview'
import React, { useEffect } from 'react'
import { SearchResultType, SearchSortBy } from 'types/search'

export interface StateToProps {
  isUserSignedIn: boolean
  searchBarText: string
  query: string
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
  loadPlaybacks: (episodeIds: string[]) => void
}

const SearchResultsList: React.FC<StateToProps & DispatchToProps> = ({
  isUserSignedIn,
  searchBarText,
  query,
  resultType,
  sortBy,
  podcastIds,
  episodeIds,
  receivedAll,
  isLoadingMore,
  loadMore,
  loadPlaybacks,
}) => {
  useEffect(() => {
    if (resultType === 'episode') {
      loadPlaybacks(episodeIds)
    }
  }, [])

  useEffect(() => {
    if (resultType === 'episode') {
      loadPlaybacks(episodeIds)
    }
  }, [isUserSignedIn, query, resultType, sortBy])

  if (resultType === 'podcast') {
    if (podcastIds.length === 0) {
      return isLoadingMore ? (
        <></>
      ) : (
        <div className="ml-5 text-gray-800 tracking-wider">
          {'No results found'}
        </div>
      )
    }

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
    if (episodeIds.length === 0) {
      return isLoadingMore ? (
        <></>
      ) : (
        <div className="ml-5 text-gray-800 tracking-wider">
          {'No results found'}
        </div>
      )
    }

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
