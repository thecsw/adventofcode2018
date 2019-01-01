TEST = ARGV.delete('-t')
VERBOSE = ARGV.delete('-v')

DIST = TEST ? 32 : 10000

input = if TEST
  <<~A.strip
  1, 1
  1, 6
  8, 3
  3, 4
  5, 5
  8, 9
  A
else
  ARGV.empty? ? DATA : ARGF
end

points = input.each_line.map { |l|
  l.split(', ').map(&:to_i).freeze
}.freeze

ymin, ymax = points.map(&:first).minmax
xmin, xmax = points.map(&:last).minmax
height = ymax - ymin + 1
width = xmax - xmin + 1

# Part 1
# Simultaneous flood-fill out from all labeled points.

# Compress all points into a single integer,
# rather than creating two-element [y, x] arrays.
# Creating tons of arrays takes unreasonably long.
flat_points = points.map { |(y, x)| (y - ymin) * width + (x - xmin) }.freeze

owned = [1] * points.size
infinite = [false] * points.size

queues = flat_points.map { |p| [p] }

seen_dist = [nil] * (height * width)
claim = [nil] * (height * width)
flat_points.each_with_index { |p, i|
  claim[p] = i
  seen_dist[p] = 0
}

dist = 0

until queues.all?(&:empty?)
  dist += 1

  queues = queues.map.with_index { |q, i|
    q.flat_map { |p|
      nq = []
      # Note that any point on the bounding box causes its owner to be infinite,
      # since a step away from the bounding box increases all distances.
      # Unfortunately, we do need to unflatten here,
      # because we need to check individual coordinates.
      y, x = p.divmod(width)
      if y == 0;          infinite[i] = true else nq << p - width end
      if y == height - 1; infinite[i] = true else nq << p + width end
      if x == 0;          infinite[i] = true else nq << p - 1 end
      if x == width - 1;  infinite[i] = true else nq << p + 1 end
      # Exclude points that have been claimed in an earlier iteration.
      # (Could do this check above, but code would be repeated 4x)
      nq.reject! { |pp| seen_dist[pp] &.< dist }
      nq
    }.uniq.each { |p|
      # Either claim an unclaimed point or clash with a claimant.
      claim[p] = claim[p] ? :clash : i
    }
  }

  queues.each_with_index { |nq, i|
    nq.reject! { |p|
      if claim[p] == i
        # Only claimant this round - claim is confirmed.
        owned[i] += 1
        seen_dist[p] = dist
      end
      # Remove points where we clashed,
      # since neither claimant will make progress.
      claim[p] == :clash
    }
  }
end

p owned.zip(infinite) if VERBOSE
puts owned.zip(infinite).reject(&:last).map(&:first).max

# Margin may be needed if points are too close together.
MARGIN = 0
yrange = (ymin - MARGIN)..(ymax + MARGIN)
xrange = (xmin - MARGIN)..(xmax + MARGIN)

within = 0

y_dists = yrange.map { |y| [y, points.sum { |yy, _| (yy - y).abs }] }.to_h.freeze
x_dists = xrange.map { |x| [x, points.sum { |_, xx| (xx - x).abs }] }.to_h.freeze

yrange.each { |y|
  edge_y = y == yrange.begin || y == yrange.end
  ydist = y_dists[y]

  xrange.each { |x|
    edge_x = x == xrange.begin || x == xrange.end
    total_dist = ydist + x_dists[x]

    if total_dist < DIST
      within += 1
      puts "DANGER! SAFE ON EDGE #{y}, #{x}" if edge_y || edge_x
    end
  }
}

puts within

__END__
158, 163
287, 68
76, 102
84, 244
162, 55
272, 335
345, 358
210, 211
343, 206
219, 323
260, 238
83, 94
137, 340
244, 172
335, 307
52, 135
312, 109
276, 93
288, 274
173, 211
125, 236
200, 217
339, 56
286, 134
310, 192
169, 192
313, 106
331, 186
40, 236
194, 122
244, 76
159, 282
161, 176
262, 279
184, 93
337, 284
346, 342
283, 90
279, 162
112, 244
49, 254
63, 176
268, 145
334, 336
278, 176
353, 135
282, 312
96, 85
90, 105
354, 312