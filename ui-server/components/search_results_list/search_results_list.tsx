import EpisodeListItem from 'components/episode_list_item'
import React from 'react'
import { SearchResultType } from 'types/search'
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
        podcastIds.map((id) => <ResultPodcastItem key={id} podcastId={id} />)}
      {resultType === 'episode' &&
        episodeIds.map((id) => <EpisodeListItem episodeId={id} key={id} />)}
    </>
  )
}

export default SearchResultsList
