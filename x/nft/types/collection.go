package types

func NewCollection(denom Denom, nfts []NFTI) (c Collection) {
	c.Denom = denom
	for _, nft := range nfts {
		c.AddNFT(nft.(NFT))
	}
	return
}

func (c *Collection) AddNFT(nft NFT) {
	c.NFTs = append(c.NFTs, nft)
}
