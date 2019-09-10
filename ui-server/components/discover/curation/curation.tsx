import ButtonWithIcon from 'components/button_with_icon'
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
      <div className="flex justify-between py-6">
        <h3 className="block text-xl text-gray-800">{curation.title}</h3>
        <ButtonWithIcon icon="arrow-right" className="w-5 text-green-600" />
      </div>
      <ResponsiveGrid rows={1}>
        {podcasts.map((p) => (
          <PodcastThumbnail podcast={p} key={p.id} />
        ))}
      </ResponsiveGrid>
    </div>
  )
}

export default CurationView
