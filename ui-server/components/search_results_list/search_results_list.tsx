import React from 'react'

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
        <h1 key={id}>{id}</h1>
      ))}
    </>
  )
}

export default SearchResultsList
