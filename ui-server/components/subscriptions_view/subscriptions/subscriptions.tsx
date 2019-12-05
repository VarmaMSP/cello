import Grid from 'components/grid'
import { PodcastLink } from 'components/link'
import React, { useState } from 'react'
import { Podcast } from 'types/app'
import { getImageUrl } from 'utils/dom'
export interface StateToProps {
  subscriptions: Podcast[]
}

const Subscriptions: React.FC<StateToProps> = ({ subscriptions }) => {
  const [showAll, setShowAll] = useState(false)

  console.log('rendering')
  return (
    <div className="bg-gray-200 py-3 px-2 rounded-xl">
      <h2 className="text-lg text-gray-700 mb-4 px-2">
        {"You're subscribed to"}
      </h2>
      <Grid
        rows={showAll ? undefined : 3}
        cols={{ LG: 4, MD: 4, SM: 4 }}
        totalRowSpacing={{ LG: 10, MD: 10, SM: 10 }}
        className="md:mb-4 mb-2"
      >
        {subscriptions.map((podcast) => (
          <PodcastLink podcastId={podcast.id} key={podcast.id}>
            <a>
              <img
                className="w-full h-auto flex-none object-contain rounded-lg border cursor-pointer"
                src={getImageUrl(podcast.urlParam)}
              />
            </a>
          </PodcastLink>
        ))}
      </Grid>
      {!showAll && (
        <>
          <hr className="my-1" />
          <button
            className="w-full text-center text-gray-700"
            onClick={(e) => {
              e.preventDefault()
              setShowAll(true)
            }}
          >
            SHOW ALL
          </button>
        </>
      )}
    </div>
  )
}

export default Subscriptions
