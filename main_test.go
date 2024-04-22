package main

import (
	"testing"
	"time"
)

type MockRepository struct {
	catatan []CatatanHarian
}

func (m *MockRepository) TambahCatatan(catatan CatatanHarian) error {
	m.catatan = append(m.catatan, catatan)
	return nil
}

func (m *MockRepository) CariCatatan(tanggal time.Time) ([]CatatanHarian, error) {
	var hasil []CatatanHarian
	for _, c := range m.catatan {
		if c.Tanggal.Equal(tanggal) {
			hasil = append(hasil, c)
		}
	}
	return hasil, nil
}

func (m *MockRepository) HapusCatatan(id int) error {
	for i, c := range m.catatan {
		if c.ID == id {
			m.catatan = append(m.catatan[:i], m.catatan[i+1:]...)
			return nil
		}
	}
	return nil
}

func TestTambahCatatan(t *testing.T) {
	mockRepo := &MockRepository{}
	service := NewService(mockRepo)

	catatanBaru := CatatanHarian{ID: 1, Tanggal: time.Now(), IsiCatatan: "Catatan test"}
	err := service.TambahCatatan(catatanBaru)

	if err != nil {
		t.Errorf("Gagal menambahkan catatan: %v", err)
	}
}

func TestCariCatatan(t *testing.T) {
	mockRepo := &MockRepository{}
	service := NewService(mockRepo)

	catatanHariIni := CatatanHarian{ID: 1, Tanggal: time.Now(), IsiCatatan: "Catatan hari ini"}
	mockRepo.TambahCatatan(catatanHariIni)

	hasil, err := service.CariCatatan(time.Now())

	if err != nil {
		t.Errorf("Gagal mencari catatan: %v", err)
	}

	if len(hasil) != 1 {
		t.Errorf("Jumlah catatan tidak sesuai. Seharusnya 1, got: %d", len(hasil))
	}
}

func TestHapusCatatan(t *testing.T) {
	mockRepo := &MockRepository{}
	service := NewService(mockRepo)

	catatan := CatatanHarian{ID: 1, Tanggal: time.Now(), IsiCatatan: "Catatan untuk dihapus"}
	mockRepo.TambahCatatan(catatan)

	err := service.HapusCatatan(1)

	if err != nil {
		t.Errorf("Gagal menghapus catatan: %v", err)
	}

	hasil, _ := mockRepo.CariCatatan(time.Now())
	if len(hasil) != 0 {
		t.Errorf("Catatan tidak terhapus. Masih ada %d catatan", len(hasil))
	}
}
