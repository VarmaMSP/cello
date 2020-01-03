import React from 'react'
import ResultPodcastItem from './result_podcast_item'

export interface StateToProps {
  podcastIds: string[]
  receivedAll: boolean
}

export interface OwnProps {
  searchQuery: string
}

const SearchResultsList: React.FC<StateToProps & OwnProps> = ({
  podcastIds,
  searchQuery,
}) => {
  return (
    <>
      <div className="-mt-1 mb-5 text-gray-700 text-lg lg:text-xl">{`Podcasts matching "${searchQuery}"`}</div>
      {podcastIds.map((id) => (
        <ResultPodcastItem key={id} podcastId={id} searchQuery={searchQuery}/>
      ))}
    </>
  )
}

export default SearchResultsList
