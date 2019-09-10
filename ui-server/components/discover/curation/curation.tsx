import PodcastThumbnail from 'components/podcast_thumbnail'
import ResponsiveGrid from 'components/responsive_grid'
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
      <div className="py-5">
        <h3 className="text-xl text-gray-800">{curation.title}</h3>
      </div>
      <ResponsiveGrid
        rows={1}
        cols={7}
        className="md:w-36 w-32 md:mx-0 mx-2"
        rowSpacing={4}
      >
        {podcasts.map((p) => (
          <PodcastThumbnail podcast={p} key={p.id} />
        ))}
      </ResponsiveGrid>
    </div>
  )
}

export default CurationView
