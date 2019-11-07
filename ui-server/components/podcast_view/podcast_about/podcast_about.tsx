import React from 'react'
import { Podcast } from 'types/app'

export interface StateToProps {
  podcast: Podcast
}

export interface OwnProps {
  podcastId: string
}

interface Props extends StateToProps, OwnProps {}

const PodcastAbout: React.FC<Props> = ({ podcast }) => {
  return (
    <div
      className="mt-5 text-gray-800 leading-snug"
      style={{ hyphens: 'auto' }}
    >
      {podcast.description}
    </div>
  )
}

export default PodcastAbout
