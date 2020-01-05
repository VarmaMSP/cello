import React from 'react'
import ResultEpisodeItem from './result_episode_item'
import ResultPodcastItem from './result_podcast_item'

export interface StateToProps {
  podcastIds: string[]
  episodeIds: string[]
  receivedAll: boolean
}

export interface OwnProps {
  searchQuery: string
  resultType: 'podcast' | 'episode'
  sortBy: 'relevance' | 'publish_date'
}

const SearchResultsList: React.FC<StateToProps & OwnProps> = ({
  podcastIds,
  episodeIds,
  searchQuery,
  resultType,
}) => {
  return (
    <>
      {resultType === 'podcast' &&
        podcastIds.map((id) => (
          <ResultPodcastItem
            key={id}
            podcastId={id}
            searchQuery={searchQuery}
          />
        ))}
      {resultType === 'episode' &&
        episodeIds.map((id) => (
          <ResultEpisodeItem
            key={id}
            episodeId={id}
            searchQuery={searchQuery}
          />
        ))}
    </>
  )
}

export default SearchResultsList
