import { imageUrl } from 'components/utils'
import React from 'react'
import { Podcast } from 'types/app'
import SubscribeButton from './components/subscribe_button'

export interface StateToProps {
  podcast: Podcast
}

export interface OwnProps {
  podcastId: string
}

interface Props extends StateToProps, OwnProps {}

const PodcastDetails: React.SFC<Props> = ({ podcast }) => {
  return (
    <div className="flex">
      <img
        className="lg:h-56 h-36 lg:w-56 w-36 flex-none object-contain object-center rounded-lg border"
        src={imageUrl(podcast.id, 'md')}
      />
      <div className="flex flex-col flex-auto w-1/2 justify-between lg:px-5 px-3">
        <div className="w-full">
          <h2 className="md:text-2xl text-lg text-gray-900 leading-tight line-clamp-2">
            {podcast.title}
          </h2>
          <h3 className="md:text-base text-sm text-gray-800 leading-loose truncate">
            {podcast.author}
          </h3>
          <p className="hidden mt-1 text-sm leading-snug text-gray-600 lg:line-clamp-3">
            {podcast.description}
          </p>
        </div>
        <SubscribeButton
          className="md:w-32 w-24 md:h-10 h-8 md:text-base text-sm"
          podcastId={podcast.id}
        />
      </div>
    </div>
  )
}

export default PodcastDetails
