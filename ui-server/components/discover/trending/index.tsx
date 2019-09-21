import Grid from 'components/grid'
import PodcastThumbnail from 'components/podcast_thumbnail'
import React from 'react'
import { connect } from 'react-redux'
import { makeGetTrendingPodcasts } from 'selectors/entities/podcasts'
import { AppState } from 'store'
import { Podcast } from 'types/app'

export interface StateToProps {
  trendingPodcasts: Podcast[]
}

const Trending: React.SFC<StateToProps> = (props) => {
  const { trendingPodcasts } = props
  return (
    <Grid rows={3} className="mb-3">
      {trendingPodcasts.map((p) => (
        <PodcastThumbnail podcast={p} />
      ))}
    </Grid>
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
