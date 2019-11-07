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
    <div className="text-black tracking-wide leading-snug">
      {podcast.description}
    </div>
  )
}

export default PodcastAbout
