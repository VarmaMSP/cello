import Grid from 'components/grid'
import Router from 'next/router'
import React from 'react'
import { Podcast } from 'types/app'
import { getImageUrl } from 'utils/dom'

export interface StateToProps {
  trendingPodcasts: Podcast[]
}

const ListTrendingPodcasts: React.SFC<StateToProps> = (props) => {
  const { trendingPodcasts } = props
  const onPodcastSelect = (podcastId: string) => () =>
    Router.push(
      {
        pathname: '/podcasts',
        query: { podcastId },
      },
      `/podcasts/${podcastId}`,
    )

  return (
    <>
      <h3 className="text-xl font-sans">{'Podcasts trending today'}</h3>
      <Grid
        cols={{ LG: 3, MD: 1, SM: 1 }}
        totalRowSpacing={{ LG: 2, MD: 10, SM: 0 }}
        className="my-5"
        classNameChild="flex md:hover:bg-gray-200 rounded-xl cursor-pointer"
      >
        {trendingPodcasts.map((podcast) => (
          <div
            className="flex w-full h-full md:p-3"
            onClick={onPodcastSelect(podcast.id)}
          >
            <img
              className="md:w-28 w-30 md:h-28 h-30 flex-none object-contain rounded-lg border"
              src={getImageUrl(podcast.id, 'md')}
            />
            <div className="w-2/3 pl-3">
              <div className="w-full font-sans text-gray-800 tracking-tight mb-2 line-clamp-2">
                {podcast.title}
              </div>
              <p className="text-sm text-gray-600 leading-snug tracking-tight line-clamp-3">
                {podcast.description}
              </p>
            </div>
          </div>
        ))}
      </Grid>
    </>
  )
}

export default ListTrendingPodcasts
