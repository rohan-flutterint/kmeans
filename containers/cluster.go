package containers

import (
	"fmt"
	"math"
)

type Cluster struct {
	center  Vector
	members []Vector
}

func NewCluster(center Vector) Cluster {
	return Cluster{
		center: center,
	}
}

func (c *Cluster) Recenter() error {
	memberCnt := len(c.members)
	if memberCnt == 0 {
		return nil
		//return fmt.Errorf("there is no mean for an empty cluster")
	}

	// newCenter = "Mean" of the Members
	newCenter := make(Vector, len(c.center))
	for _, member := range c.members {
		newCenter.Add(member)
	}
	newCenter.Mul(1 / float64(memberCnt))

	c.center = newCenter
	return nil
}

func (c *Cluster) RecenterReturningMovedDistance(distFn DistanceFunction) (moveDistances float64, err error) {
	memberCnt := len(c.members)
	if memberCnt == 0 {
		//return 0, errors.New("kmeans: there is no mean for an empty cluster")
		return 0, nil
	}

	// newCenter is the "Mean" of the Members
	newCenter := make(Vector, len(c.center))
	for _, member := range c.members {
		newCenter.Add(member)
	}
	newCenter.Mul(1 / float64(memberCnt))

	moveDistances, err = distFn(c.center, newCenter)
	if err != nil {
		return 0, err
	}

	c.center = newCenter

	return moveDistances, nil
}

// SSE returns the sum of squared errors of the cluster
func (c *Cluster) SSE() float64 {
	sse := 0.0
	for i := 0; i < len(c.members); i++ {
		dist, _ := EuclideanDistance(c.center, c.members[i])
		sse += math.Pow(dist, 2)
	}
	return sse
}

// Reset only resets the members of the cluster. The center is not reset.
func (c *Cluster) Reset() {
	c.members = []Vector{}
}

func (c *Cluster) String() string {
	return fmt.Sprintf("Center: %v, Members: %v", c.center, c.members)
}

func (c *Cluster) GetCenter() Vector {
	return c.center
}

func (c *Cluster) GetMembers() []Vector {
	return c.members
}

func (c *Cluster) AddMember(member Vector) {
	c.members = append(c.members, member)
}
