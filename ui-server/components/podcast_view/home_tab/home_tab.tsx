import React from 'react'
import { Podcast } from 'types/app'
import EpisodeList from '../episode_list'

export interface StateToProps {
  podcast: Podcast
}

export interface OwnProps {
  podcastId: string
}

interface Props extends StateToProps, OwnProps {}

const PodcastAbout: React.FC<Props> = ({ podcast }) => {
  return (
    <div>
      <h2 className="font-medium tracking-wider mb-2">{'Description'}</h2>
      <div
        className="text-gray-800 text-sm tracking-wide leading-relaxed"
        style={{ hyphens: 'auto' }}
      >
        <div>{podcast.description}</div>
        <div className="mt-5 text-gray-600">{`${podcast.copyright}`}</div>
      </div>

      <hr className="my-6" />

      <h2 className="font-medium tracking-widest mb-5">{'Episodes'}</h2>
      <EpisodeList podcastId={podcast.id} />
    </div>
  )
}

export default PodcastAbout
