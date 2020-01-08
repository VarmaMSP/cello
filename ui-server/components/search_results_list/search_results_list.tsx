import EpisodeListItem from 'components/episode_list_item'
import PodcastPreview from 'components/podcast_preview'
import React from 'react'
import { SearchResultType } from 'types/search'

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
        podcastIds.map((id) => <PodcastPreview key={id} podcastId={id} />)}
      {resultType === 'episode' &&
        episodeIds.map((id) => <EpisodeListItem episodeId={id} key={id} />)}
    </>
  )
}

export default SearchResultsList
