import React from 'react'
import { Episode, Podcast } from 'types/app'

export interface StateToProps {
  episode: Episode
  podcast: Podcast
}

export interface OwnProps {
  episodeId: string
}

const EpisodeView: React.FC<StateToProps & OwnProps> = ({
  episode,
  podcast,
}) => {
  return (
    <div>
      <h1>{episode.title}</h1>
      <h2>{podcast.title}</h2>
    </div>
  )
}

export default EpisodeView
