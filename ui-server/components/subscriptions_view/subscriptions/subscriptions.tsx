import Grid from 'components/grid'
import { PodcastLink } from 'components/link'
import React from 'react'
import { Podcast } from 'types/app'
import { getImageUrl } from 'utils/dom'
export interface StateToProps {
  subscriptions: Podcast[]
}

const Subscriptions: React.FC<StateToProps> = ({ subscriptions }) => {
  return (
    <div className="pl-8 pr-2 rounded-xl">
      <h2 className="text-lg text-gray-700">{"You're subscribed to"}</h2>
      <hr className="my-4 border-gray-400" />
      <Grid
        cols={{ LG: 4, MD: 4, SM: 4 }}
        totalRowSpacing={{ LG: 10, MD: 10, SM: 10 }}
        className="md:mb-4 mb-2"
      >
        {subscriptions.map((podcast) => (
          <PodcastLink podcastUrlParam={podcast.urlParam} key={podcast.id}>
            <a>
              <img
                className="w-full h-auto flex-none object-contain rounded-lg border cursor-pointer"
                src={getImageUrl(podcast.urlParam)}
              />
            </a>
          </PodcastLink>
        ))}
      </Grid>
    </div>
  )
}

export default Subscriptions
