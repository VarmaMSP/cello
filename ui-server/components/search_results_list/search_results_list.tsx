import React from 'react'
import { SearchResultType } from 'types/search'
import ResultEpisodeItem from './result_episode_item'
import ResultPodcastItem from './result_podcast_item'

export interface StateToProps {
  resultType: SearchResultType
  podcastIds: string[]
  episodeIds: string[]
  receivedAll: boolean
}

const SearchResultsList: React.FC<StateToProps> = ({
  podcastIds,
  episodeIds,
  resultType,
}) => {
  return (
    <>
      {resultType === 'podcast' &&
        podcastIds.map((id) => (
          <ResultPodcastItem
            key={id}
            podcastId={id}
          />
        ))}
      {resultType === 'episode' &&
        episodeIds.map((id) => (
          <ResultEpisodeItem
            key={id}
            episodeId={id}
          />
        ))}
    </>
  )
}

export default SearchResultsList
