import React from 'react'
import { Podcast } from '../../types/app'

export interface StateToProps {
  podcast: Podcast
}

export interface OwnProps {
  podcastId: string
}

interface Props extends StateToProps, OwnProps {}

const PodcastDetails: React.SFC<Props> = ({ podcast }) => (
  <div className="flex mb-8">
    <img
      className="lg:h-56 lg:w-56 h-32 w-32 flex-none object-cover object-center rounded"
      src={`http://localhost:8080/img/${podcast.id}p-500x500.jpg`}
    />
    <div className="flex flex-col lg:px-5 px-3">
      <h2 className="text-2xl text-gray-900">{podcast.title}</h2>
      <h3 className="text-lg text-gray-800 leading-relaxed">
        {podcast.author}
      </h3>
      <span className="lg:block hidden mt-1 text-sm text-gray-800 ">
        {podcast.description}
      </span>
    </div>
  </div>
)

export default PodcastDetails
