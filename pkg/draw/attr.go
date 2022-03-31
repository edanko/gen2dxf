package draw

func (d *Drawer) addAttrs(quality, quantity, thickness, order, section, partid, ina string) error {
	as := []string{
		"[M]" + quality,
		"[Q]" + quantity,
		"[T]" + thickness,
		"[B]" + order,
		"[S]" + section,
		"[P]" + partid,
		"[I]" + ina,
		"[N]" + ina,
		"[A]" + ina,
	}

	for i, a := range as {
		t, err := d.drawing.Text(a, 0, -float64(i)*attrSpacing, 0, attrTextHeight)
		if err != nil {
			return err
		}
		t.SetColor(attrColor)
	}
	return nil
}
