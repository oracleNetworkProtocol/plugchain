package types

func NewCollection(denom Class, nfts []NFTI) (c Collection) {
	c.Class = denom
	for _, nft := range nfts {
		c.AddNFT(nft.(NFT))
	}
	return
}

func (c *Collection) AddNFT(nft NFT) {
	c.NFTs = append(c.NFTs, nft)
}
