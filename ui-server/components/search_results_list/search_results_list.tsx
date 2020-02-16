import ButtonShowMore from 'components/button_show_more'
import EpisodePreview from 'components/episode_preview'
import { PodcastLink } from 'components/link'
import PodcastPreview from 'components/podcast_preview'
import { Podcast } from 'models'
import React, { useEffect } from 'react'
import { SearchResultType, SearchSortBy } from 'types/search'
import { getImageUrl } from 'utils/dom'

export interface StateToProps {
  isUserSignedIn: boolean
  searchBarText: string
  query: string
  resultType: SearchResultType
  sortBy: SearchSortBy
  podcastIds: string[]
  podcastsBestMatch: Podcast[]
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
  podcastsBestMatch,
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
        <div className="mt-6 text-center text-gray-800 tracking-wider">
          {'No results found'}
        </div>
      )
    }

    return (
      <div>
        {podcastIds.map((id) => (
          <div key={id} className="mb-6">
            <PodcastPreview podcastId={id} showHighlights />
          </div>
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
        <div className="mt-6 text-center text-gray-800 tracking-wider">
          {'No results found'}
        </div>
      )
    }

    return (
      <div>
        {podcastsBestMatch.length > 0 && (
          <>
            <div className="text-lg font-semibold my-3">Podcasts</div>
            <div className="flex">
              {podcastsBestMatch.map((p) => (
                <div
                  key={p.id}
                  className="flex-none md:w-28 w-22 mx-4 mt-3 mb-6"
                >
                  <PodcastLink podcastUrlParam={p.urlParam}>
                    <a>
                      <img
                        className="w-full h-auto mb-2 flex-none object-contain rounded-lg border"
                        src={getImageUrl(p.urlParam)}
                      />
                    </a>
                  </PodcastLink>
                  <PodcastLink podcastUrlParam={p.urlParam}>
                    <a className="text-sm tracking-wide font-medium leading-tight line-clamp-2">
                      {p.title}
                    </a>
                  </PodcastLink>
                  <PodcastLink podcastUrlParam={p.urlParam}>
                    <a className="text-2xs text-gray-900 tracking-wide font-medium leading-loose line-clamp-1">
                      {p.author}
                    </a>
                  </PodcastLink>
                </div>
              ))}
            </div>
          </>
        )}

        <div className="text-lg font-semibold my-3">Episodes</div>
        {episodeIds.map((id) => (
          <div key={id} className="mb-6">
            <EpisodePreview episodeId={id} showHighlights />
          </div>
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
