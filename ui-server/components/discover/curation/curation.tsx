import Grid from 'components/grid'
import PodcastThumbnail from 'components/podcast_thumbnail'
import React from 'react'
import { Curation, Podcast } from 'types/app'

export interface StateToProps {
  curation: Curation
  podcasts: Podcast[]
}

export interface OwnProps {
  curationId: string
}

interface Props extends StateToProps, OwnProps {}

const CurationView: React.SFC<Props> = (props) => {
  const { curation, podcasts } = props

  return (
    <div>
      <div className="pb-5">
        <h3 className="text-xl text-gray-800">{curation.title}</h3>
      </div>
      <Grid rows={1} cols={7} className="mb-4 pb-8" classNameChild="w-36 mx-2">
        {podcasts.map((p) => (
          <PodcastThumbnail podcast={p} key={p.id} />
        ))}
      </Grid>
    </div>
  )
}

export default CurationView
