import { PodcastLink } from 'components/link'
import React from 'react'
import { Podcast } from 'types/models'
import { getImageUrl } from 'utils/dom'

export interface StateToProps {
  subscriptions: Podcast[]
}

const SubscriptionsList: React.FC<StateToProps> = ({ subscriptions }) => {
  if (subscriptions.length === 0) {
    return <></>
  }

  return (
    <div className="px-3 pt-3 pb-2 rounded-xl bg-gray-200">
      <h2 className="mb-3 text text-gray-800 tracking-wide">
        {"You're subscribed to"}
      </h2>
      <div className="flex flex-wrap">
        {subscriptions.map((podcast) => (
          <PodcastLink podcastUrlParam={podcast.urlParam} key={podcast.id}>
            <a
              className="block mb-4 lg:w-1/5 md:w-1/4 w-1/5"
              style={{ paddingLeft: '2px', paddingRight: '2px' }}
            >
              <img
                className="w-full h-auto flex-none object-contain rounded-lg border cursor-pointer"
                src={getImageUrl(podcast.urlParam)}
              />
            </a>
          </PodcastLink>
        ))}
      </div>
    </div>
  )
}

export default SubscriptionsList
