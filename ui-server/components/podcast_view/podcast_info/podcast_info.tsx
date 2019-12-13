import ButtonSubscribe from 'components/button_subscribe'
import format from 'date-fns/format'
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

const PodcastInfo: React.SFC<Props> = ({ podcast }) => {
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
          images: [{ url: getImageUrl(podcast.urlParam) }],
        }}
        twitter={{
          cardType: `summary_large_image`,
        }}
      />
      <div className="flex">
        <img
          className="lg:h-36 h-24 lg:w-36 w-24 flex-none object-contain object-center rounded-lg border"
          src={getImageUrl(podcast.urlParam)}
        />
        <div className="flex flex-col flex-auto w-1/2 justify-between lg:px-5 px-3">
          <div className="w-full">
            <h2 className="md:text-xl text-lg text-gray-900 leading-tight line-clamp-2">
              {podcast.title}
            </h2>
            <h3 className="md:text-base text-sm text-gray-800 leading-loose truncate">
              {podcast.author}
            </h3>
            <h4 className="text-xs text-gray-700">
              {`${podcast.totalEpisodes} episodes`}
              <span className="mx-2 font-extrabold">&middot;</span>
              {`Since ${format(
                new Date(`${podcast.earliestEpisodePubDate} +0000`),
                'MMM yyyy',
              )}`}
            </h4>
          </div>
          <ButtonSubscribe
            className="w-24 py-1 text-xs"
            podcastId={podcast.id}
          />
        </div>
      </div>
    </>
  )
}

export default PodcastInfo
