import NavTabs from 'components/nav_tabs'
import { NextSeo } from 'next-seo'
import React from 'react'
import { connect } from 'react-redux'
import { getPodcastById } from 'selectors/entities/podcasts'
import { AppState } from 'store'
import { Podcast } from 'types/models'
import { getImageUrl } from 'utils/dom'
import PodcastHeader from './components/podcast_header/podcast_header'
import HomeTab from './tabs/home'

interface StateToProps {
  podcast: Podcast
}

interface OwnProps {
  podcastId: string
  activeTab?: string
}

const PodcastView: React.FC<StateToProps & OwnProps> = ({
  podcast,
  activeTab,
}) => {
  const podcastUrlParam = podcast.urlParam

  return (
    <div className="pt-6">
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

      <PodcastHeader podcast={podcast} />
      <div className="mt-6 mb-4">
        <NavTabs
          tabs={[
            {
              name: 'podcast',
              pathname: '/podcasts',
              query: { podcastUrlParam, skipLoad: true },
              as: `/podcasts/${podcastUrlParam}`,
            },
          ]}
          active={activeTab}
          defaultTab="podcast"
        />
      </div>
      {activeTab === undefined && <HomeTab podcast={podcast} />}
    </div>
  )
}

function mapStateToProps(
  state: AppState,
  { podcastId }: OwnProps,
): StateToProps {
  return {
    podcast: getPodcastById(state, podcastId),
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(mapStateToProps)(
  PodcastView,
)
