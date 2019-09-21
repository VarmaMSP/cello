import { imageUrl } from 'components/utils'
import Router from 'next/router'
import React from 'react'
import Shiitake from 'shiitake'
import { Podcast } from 'types/app'

export interface StateToProps {
  trendingPodcasts: Podcast[]
}

const TrendingPodcastsList: React.SFC<StateToProps> = (props) => {
  const { trendingPodcasts } = props
  const onPodcastSelect = (podcastId: string) => () =>
    Router.push(
      {
        pathname: '/podcasts',
        query: { id: podcastId },
      },
      `/podcasts/${podcastId}`,
    )

  let colsJsx: JSX.Element[] = []
  for (let i = 0; i < trendingPodcasts.length; ++i) {
    const podcast = trendingPodcasts[i]
    colsJsx.push(
      <div
        className="flex md:w-6/13 w-full my-5 md:p-3 lg:hover:bg-gray-200 cursor-pointer rounded-xl"
        onClick={onPodcastSelect(podcast.id)}
      >
        <img
          className="md:w-32 w-30 md:h-32 h-30 flex-none object-contain rounded-lg border"
          src={imageUrl(podcast.id, 'md')}
        />
        <div className="w-2/3 py-1 pl-3">
          <div className="w-full font-sans text-gray-800 tracking-tight mb-2">
            {podcast.title}
          </div>
          <Shiitake
            lines={4}
            throttleRate={200}
            className="text-sm text-gray-600 leading-snug tracking-tight"
          >
            {podcast.description}
          </Shiitake>
        </div>
      </div>,
    )
  }
  if (colsJsx.length % 2 != 0) {
    colsJsx.push(<div className="md:w-6/13 w-full" />)
  }

  let rowsJsx: JSX.Element[] = []
  for (let i = 0; i < colsJsx.length; i += 2) {
    rowsJsx.push(
      <div className="flex md:flex-row flex-col justify-around">
        {colsJsx[i]}
        {colsJsx[i + 1]}
      </div>,
    )
  }

  return (
    <>
      <h3 className="text-xl font-sans">{'Podcasts trending today'}</h3>
      {rowsJsx}
    </>
  )
}

export default TrendingPodcastsList
