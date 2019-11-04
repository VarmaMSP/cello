import ButtonSubscribe from 'components/button_subscribe'
import { NextSeo } from 'next-seo'
import React from 'react'
import { Podcast } from 'types/app'
import { getImageUrl } from 'utils/dom'

export interface StateToProps {
  podcast: Podcast
}

export interface OwnProps {
  podcastId: string
}

interface Props extends StateToProps, OwnProps {}

const PodcastDetails: React.SFC<Props> = ({ podcast }) => {
  return (
    <>
      <NextSeo
        title={`${podcast.title} | Phenopod`}
        description={podcast.description}
        canonical={`https://phenopod.com/podcasts/${podcast.id}`}
        openGraph={{
          url: `https://phenopod.com/podcasts/${podcast.id}`,
          type: 'article',
          title: podcast.title,
          description: podcast.description,
          images: [{ url: getImageUrl(podcast.id, 'md') }],
        }}
        twitter={{
          cardType: `summary_large_image`,
        }}
      />
      <div className="flex">
        <img
          className="lg:h-36 h-24 lg:w-36 w-24 flex-none object-contain object-center rounded-lg border"
          src={getImageUrl(podcast.id, 'md')}
        />
        <div className="flex flex-col flex-auto w-1/2 justify-between lg:px-5 px-3">
          <div className="w-full">
            <h2 className="md:text-2xl text-lg text-gray-900 leading-tight line-clamp-2">
              {podcast.title}
            </h2>
            <h3 className="md:text-base text-sm text-gray-800 leading-loose truncate">
              {podcast.author}
            </h3>
          </div>
          <ButtonSubscribe
            className="w-28 px-4 py-2 text-sm"
            podcastId={podcast.id}
          />
        </div>
      </div>
    </>
  )
}

export default PodcastDetails
