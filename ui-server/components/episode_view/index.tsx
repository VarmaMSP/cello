import NavTabs from 'components/nav_tabs'
import { NextSeo } from 'next-seo'
import React from 'react'
import { connect } from 'react-redux'
import { getEpisodeById } from 'selectors/entities/episodes'
import { AppState } from 'store'
import { Episode } from 'types/app'
import EpisodeHeader from './components/episode_header'
import HomeTab from './tabs/home'

export interface StateToProps {
  episode: Episode
}

export interface OwnProps {
  episodeId: string
  activeTab: string | undefined
}

const EpisodeView: React.FC<StateToProps & OwnProps> = ({
  episode,
  activeTab,
}) => {
  const episodeUrlParam = episode.urlParam

  return (
    <div className="pt-6">
      <NextSeo
        title={`${episode.title} | Phenopod`}
        description={episode.summary}
        canonical={`https://phenopod.com/episodes/${episode.urlParam}`}
        openGraph={{
          url: `https://phenopod.com/episodes/${episode.urlParam}`,
          type: 'article',
          title: episode.title,
          description: episode.summary,
        }}
        twitter={{
          cardType: `summary_large_image`,
        }}
      />
      <EpisodeHeader episode={episode} />
      <div className="mt-6 mb-4">
        <NavTabs
          tabs={[
            {
              name: 'episode',
              pathname: '/episodes',
              query: { episodeUrlParam, skipLoad: true },
              as: `/episodes/${episodeUrlParam}`,
            },
          ]}
          active={activeTab}
          defaultTab="episode"
        />
      </div>
      <div className="mt-6 mb-4">
        <HomeTab episode={episode} />
      </div>
    </div>
  )
}

function mapStateToProps(
  state: AppState,
  { episodeId }: OwnProps,
): StateToProps {
  return { episode: getEpisodeById(state, episodeId) }
}

export default connect<StateToProps, {}, OwnProps, AppState>(mapStateToProps)(
  EpisodeView,
)
