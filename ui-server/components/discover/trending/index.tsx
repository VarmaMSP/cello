import Grid from 'components/grid'
import Link from 'next/link'
import React from 'react'
import { connect } from 'react-redux'
import { makeGetTrendingPodcasts } from 'selectors/entities/podcasts'
import { AppState } from 'store'
import { Podcast } from 'types/app'
import { getImageUrl } from 'utils/dom'

export interface StateToProps {
  trendingPodcasts: Podcast[]
}

const Trending: React.SFC<StateToProps> = (props) => {
  const { trendingPodcasts } = props
  return (
    <div className="w-full pb-8">
      <div className="flex justify-between pb-4">
        <h3 className="text-lg text-bold text-gray-900">
          {'Podcasts Trending Today'}
        </h3>
        <Link href="/trending" scroll={false}>
          <a>
            <div className="flex align-center text-lg text-green-600">
              {'see more'}
            </div>
          </a>
        </Link>
      </div>
      <Grid
        rows={{ LG: 4, MD: 6, SM: 6 }}
        cols={{ LG: 8, MD: 6, SM: 4 }}
        totalRowSpacing={{ LG: 12, MD: 8, SM: 5 }}
        className="md:mb-3 mb-2"
      >
        {trendingPodcasts.map((podcast) => (
          <Link
            href={{ pathname: '/podcasts', query: { podcastId: podcast.id } }}
            as={`/podcasts/${podcast.id}`}
            key={podcast.id}
          >
            <a>
              <img
                className="w-full h-auto flex-none object-contain rounded-lg border cursor-pointer"
                src={getImageUrl(podcast.id, 'md')}
              />
            </a>
          </Link>
        ))}
      </Grid>
    </div>
  )
}

function makeMapStateToProps() {
  const getTrendingPodcasts = makeGetTrendingPodcasts()
  return (state: AppState) => ({
    trendingPodcasts: getTrendingPodcasts(state),
  })
}

export default connect<StateToProps, {}, {}, AppState>(makeMapStateToProps())(
  Trending,
)
